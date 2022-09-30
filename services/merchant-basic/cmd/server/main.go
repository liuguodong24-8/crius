package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/config"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/rpc"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	mutil "gitlab.omytech.com.cn/micro-service/merchant-basic/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GitCommit git build commit
var GitCommit string

func init() {
	pkgs.PrintGitCommit(GitCommit)
	usage()
	var conf string
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.Parse()
	getConfigPath(&conf)
	config.Load(conf)

	if err := model.DatabaseConnection(); err != nil {
		panic(fmt.Sprintf("数据库连接失败:%s", err.Error()))
	}

	if err := util.RegisterLogger(logger.Config{
		Channel:    config.Setting.Log.Channel,
		Level:      logger.Level(config.Setting.Log.Level),
		OutputFile: config.Setting.Log.Output,
		WithStack:  config.Setting.Log.Stack,
	}); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}

	if err := mutil.MQTTConnect(); err != nil {
		panic(fmt.Sprintf("MQTT连接失败:%s", err.Error()))
	}
}

func main() {
	defer util.CatchException()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Setting.App.IP, config.Setting.App.Port))
	if err != nil {
		panic(fmt.Sprintf("开启grpc端口监听失败:%s", err.Error()))
	}

	basic, err := rpc.NewMerchantBasicServe()
	if err != nil {
		panic(fmt.Sprintf("开启grpc端口监听失败:%s", err.Error()))
	}

	s := grpc.NewServer()
	proto.RegisterMerchantBasicServiceServer(s, basic)
	reflection.Register(s)

	go heart(basic)

	if err := s.Serve(listen); nil != err {
		panic(fmt.Sprintf("开启grpc失败:%s", err.Error()))
	}
}

func heart(s *rpc.Server) {
	for true {
		if err := s.Heart(context.Background()); nil != err {
			util.Logger.Error(fmt.Sprintf("basic 服务心跳失败:%s", err.Error()))
		}

		time.Sleep(time.Second * 3)
	}
}

// usage 返回使用方法
func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: go run cmd/server/main.go -c {配置文件}
`)
}

func getConfigPath(conf *string) {
	if *conf != "" {
		return
	}
	path, err := os.Getwd()
	if err != nil {
		return
	}
	*conf = fmt.Sprintf("%s/config/config.toml", path)
}
