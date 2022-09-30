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

	"gitlab.omytech.com.cn/micro-service/Crius/internal/config"
	"gitlab.omytech.com.cn/micro-service/Crius/internal/rpc"
	"gitlab.omytech.com.cn/micro-service/Crius/internal/web/router"
	"gitlab.omytech.com.cn/micro-service/Crius/proto"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
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
	// 捕获异常信息
	defer util.CatchException()

	startGrpc()
	//go startGrpc()
	//startWeb()
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
	listen, err := net.Listen("tcp", config.Setting.Grpc.Address)
	if err != nil {
		panic(fmt.Sprintf("开启grpc端口监听失败:%s", err.Error()))
	}

	crius, err := rpc.NewCriusServe()
	if err != nil {
		panic(fmt.Sprintf("开启grpc端口监听失败:%s", err.Error()))
	}

	s := grpc.NewServer()
	proto.RegisterCriusServiceServer(s, crius)
	reflection.Register(s)

	if err := s.Serve(listen); nil != err {
		panic(fmt.Sprintf("开启grpc失败:%s", err.Error()))
	}
}

// usage 返回使用方法
func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: go run cmd/server/main.go -c {配置文件}
`)
}
