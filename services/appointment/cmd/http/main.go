package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/rpc"

	_ "github.com/lib/pq"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/config"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
)

// GitCommit git build commit
var GitCommit string

func main() {
	fmt.Println("build_git_commit:", GitCommit)
	usage()
	var conf string
	var date string
	var after bool
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.StringVar(&date, "d", "", "指定日期:yyyy-MM-dd")
	flag.BoolVar(&after, "a", false, "是否发送该日期之后所有数据:true/false")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)

	if err := util.RegisterLogger(logger.Config{
		Channel:    config.Setting.Log.Channel,
		Level:      logger.Level(config.Setting.Log.Level),
		OutputFile: config.Setting.Log.Output,
		WithStack:  config.Setting.Log.Stack,
	}); nil != err {
		panic(fmt.Sprintf("注册日志失败:%s", err.Error()))
	}

	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatalf("解析日期参数错误 yyyy-MM-dd:%v", err)
	}

	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		log.Fatalf("crius client失败:%v", err)
	}

	if _, err := model.DatabaseConnection(); err != nil {
		log.Fatalf("数据库连接错误:%v", err)
	}

	appointments, err := model.GetAppointmentsBySended(t, after, false)
	if err != nil {
		log.Fatalf("查询预约数据库错误:%v", err)
	}

	s := rpc.Server{}
	s.SetCrius(client)
	for i := range appointments {
		s.AppointmentHTTPRequest(&appointments[i])
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
	Usage: go run cmd/http/main.go -c {配置文件}
	`)
}
