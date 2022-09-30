package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
	sqlMigrate "github.com/rubenv/sql-migrate"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/migrate"
	"gitlab.omytech.com.cn/micro-service/payment/internal/config"
)

// GitCommit git build commit
var GitCommit string

func main() {
	fmt.Println("build_git_commit:", GitCommit)
	migrate.Usage()
	var conf, migrateType string
	var number int
	var direction sqlMigrate.MigrationDirection
	flag.StringVar(&conf, "c", "", "指定配置文件位置")
	flag.StringVar(&migrateType, "a", "", "up 迁移数据库,down 回滚数据库")
	flag.IntVar(&number, "n", 0, "指定迁移数量")
	flag.Parse()
	if conf == "" {
		panic("请指定配置文件位置")
	}
	config.Load(conf)
	path, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "获取当前目录错误:%v", err)
		return
	}
	switch strings.ToLower(migrateType) {
	case "up":
		direction = sqlMigrate.Up
	case "down":
		direction = sqlMigrate.Down
	default:
		_, _ = fmt.Fprintf(os.Stderr, "-a 参数错误, up 迁移数据, down 回滚数据\n")
	}

	cfg := migrate.Config{
		DBName:           config.Setting.Database.DBName,
		SslModel:         config.Setting.Database.SslModel,
		User:             config.Setting.Database.User,
		Password:         config.Setting.Database.Password,
		Host:             config.Setting.Database.Host,
		Port:             int32(config.Setting.Database.Port),
		Dir:              fmt.Sprintf("%s/migrations/postgres", path),
		MigrateDirection: direction,
		MigrateNumber:    number,
		Schema:           config.Setting.Migrate.Schema,
		Table:            config.Setting.Migrate.Table,
	}
	migrate.Migrate(&cfg)

}
