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
	"gitlab.omytech.com.cn/micro-service/message/internal/config"
	"gitlab.omytech.com.cn/micro-service/message/internal/rpc"
	"gitlab.omytech.com.cn/micro-service/message/proto"
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
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Setting.App.IP, config.Setting.App.Port))
	if err != nil {
		panic(fmt.Sprintf("开启branch grpc端口监听失败:%s", err.Error()))
	}
	server, err := rpc.NewMessageServe()
	if err != nil {
		panic(fmt.Sprintf("实例branch server失败:%s", err.Error()))
	}

	s := grpc.NewServer()
	proto.RegisterMessageServiceServer(s, server)
	reflection.Register(s)

	go heart(server)

	if err := s.Serve(listen); nil != err {
		panic(fmt.Sprintf("开启branch grpc失败:%s", err.Error()))
	}
}

func heart(s *rpc.Server) {
	for true {
		if err := s.Heart(context.Background()); nil != err {
			util.Logger.Error(fmt.Sprintf("branch 服务心跳失败:%s", err.Error()))
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
