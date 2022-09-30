package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/payment/internal/config"
	"gitlab.omytech.com.cn/micro-service/payment/internal/rpc"
	"gitlab.omytech.com.cn/micro-service/payment/internal/web/router"
	"gitlab.omytech.com.cn/micro-service/payment/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GitCommit git build commit
var GitCommit string

func init() {
	fmt.Println("build_git_commit:", GitCommit)
	usage()
	var conf string
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)

}

func main() {
	if err := util.RegisterLogger(logger.Config{
		Channel:    config.Setting.Log.Channel,
		Level:      logger.Level(config.Setting.Log.Level),
		OutputFile: config.Setting.Log.Output,
		WithStack:  config.Setting.Log.Stack,
	}); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}
	defer util.CatchException()

	go startGrpc()

	startWeb()
}

func startWeb() {
	router.Init()
	web := &http.Server{
		Addr:    config.Setting.Web.Address,
		Handler: router.Router,
	}
	go func() {
		if err := web.ListenAndServe(); nil != err && http.ErrServerClosed != err {
			panic(fmt.Sprintf("启动web服务失败:%s", err.Error()))
		}
	}()

	// 接收终端信号来关闭服务
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	util.Logger.Info("关闭web服务")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := web.Shutdown(ctx); nil != err {
		util.Logger.Panic("关闭web服务异常，强制关闭")
	}
	util.Logger.Info("关闭web服务完成")
}

func startGrpc() {
	defer util.CatchException()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Setting.App.IP, config.Setting.App.Port))
	if err != nil {
		panic(fmt.Sprintf("开启payment grpc端口监听失败:%s", err.Error()))
	}
	server, err := rpc.NewPaymentServe()
	if err != nil {
		panic(fmt.Sprintf("实例payment server失败:%s", err.Error()))
	}

	s := grpc.NewServer()
	proto.RegisterPaymentServerServer(s, server)
	reflection.Register(s)

	go heart(server)

	if err := s.Serve(listen); nil != err {
		panic(fmt.Sprintf("开启payment grpc失败:%s", err.Error()))
	}
}

func heart(s *rpc.Server) {
	for true {
		if err := s.Heart(context.Background()); nil != err {
			util.Logger.Error(fmt.Sprintf("payment 服务心跳失败:%s", err.Error()))
		}

		time.Sleep(time.Second * 3)
	}
}

// usage 返回使用方法
func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `
Usage: go run cmd/server/main.go -c {配置文件}
`)
}
