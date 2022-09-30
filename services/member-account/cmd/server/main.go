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
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/config"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/model"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/rpc"
	"gitlab.omytech.com.cn/micro-service/member-account/proto"
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

	if err := model.DatabaseConnection(); err != nil {
		panic(fmt.Sprintf("数据库连接失败:%s", err.Error()))
	}

	if err := model.RedisConnection(); err != nil {
		panic(fmt.Sprintf("redis连接失败:%s", err.Error()))
	}

	if err := cutil.RegisterLogger(logger.Config{
		Channel:    config.Setting.Log.Channel,
		Level:      logger.Level(config.Setting.Log.Level),
		OutputFile: config.Setting.Log.Output,
		WithStack:  config.Setting.Log.Stack,
	}); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}

}

func main() {
	defer cutil.CatchException()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Setting.App.IP, config.Setting.App.Port))
	if err != nil {
		panic(fmt.Sprintf("开启branch grpc端口监听失败:%s", err.Error()))
	}
	server, err := rpc.NewMemberAccountServe()
	if err != nil {
		panic(fmt.Sprintf("实例branch server失败:%s", err.Error()))
	}

	s := grpc.NewServer()
	proto.RegisterMemberAccountServerServer(s, server)
	reflection.Register(s)

	go heart(server)

	if err := s.Serve(listen); nil != err {
		panic(fmt.Sprintf("开启branch grpc失败:%s", err.Error()))
	}
}

func heart(s *rpc.Server) {
	for true {
		if err := s.Heart(context.Background()); nil != err {
			cutil.Logger.Error(fmt.Sprintf("branch 服务心跳失败:%s", err.Error()))
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
