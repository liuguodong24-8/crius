package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"gitlab.omytech.com.cn/micro-service/basic/internal/config"
	"gitlab.omytech.com.cn/micro-service/basic/internal/rpc"
	"gitlab.omytech.com.cn/micro-service/basic/proto"
	"gitlab.omytech.com.cn/micro-service/basic/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	usage()
	var conf string
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)

	if err := util.RegisterLogger(); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}
}

func main() {
	defer util.CatchException()
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Setting.App.IP, config.Setting.App.Port))
	if err != nil {
		panic(fmt.Sprintf("开启grpc端口监听失败:%s", err.Error()))
	}

	basic, err := rpc.NewBasicServe()
	if err != nil {
		panic(fmt.Sprintf("开启grpc端口监听失败:%s", err.Error()))
	}

	s := grpc.NewServer()
	proto.RegisterBasicServiceServer(s, basic)
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
