package migrate

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"strings"

	migrate "github.com/rubenv/sql-migrate"
)

// Config ...
type Config struct {
	DBName   string
	SslModel string
	User     string
	Password string
	Host     string
	Port     int32

	Dir              string
	MigrateDirection migrate.MigrationDirection
	MigrateNumber    int

	Schema string
	Table  string
}

// Migrate 数据库迁移
func Migrate(cfg *Config) {
	//指定导入sql文件目录
	migrations := &migrate.FileMigrationSource{
		Dir: cfg.Dir,
	}

	migrate.SetSchema(cfg.Schema)
	migrate.SetTable(cfg.Table)

	fmt.Fprintf(os.Stdout, "连接到数据库%s:%d 数据库名:%s 确定? (yes/no)\n", cfg.Host, cfg.Port, cfg.DBName)
	var confirm string
	_, err := fmt.Scan(&confirm)
	if err != nil && err != io.EOF {
		panic(fmt.Sprintf("输入错误:%v", err))
	}
	if strings.ToLower(confirm) != "yes" {
		fmt.Fprintf(os.Stdout, "关闭数据库迁移...\n")
		return
	}
	//建立数据库连接
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s sslmode=%s user=%s password=%s host=%s port=%d",
		cfg.DBName, cfg.SslModel, cfg.User, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		panic(fmt.Sprintf("建立数据库连接错误:%v", err))
	}
	defer db.Close()

	if cfg.MigrateDirection == migrate.Up { //当type为up时，迁移数据库
		n, err := migrate.ExecMax(db, "postgres", migrations, migrate.Up, cfg.MigrateNumber)
		if err != nil {
			fmt.Fprintf(os.Stderr, "迁移操作失败:%v\n", err)
			return
		}
		fmt.Fprintf(os.Stdout, "迁移操作成功 %d 个sql文件!\n", n)
	} else if cfg.MigrateDirection == migrate.Down { //当type为down时，回滚数据库
		n, err := migrate.ExecMax(db, "postgres", migrations, migrate.Down, cfg.MigrateNumber)
		if err != nil {
			fmt.Fprintf(os.Stderr, "回滚操作失败:%v\n", err)
			return
		}
		fmt.Fprintf(os.Stdout, "回滚操作成功 %d 个sql文件!\n", n)
	} else {
		fmt.Fprintf(os.Stderr, "-a 参数错误, up 迁移数据, down 回滚数据\n")
	}
}

// Usage 返回使用方法
func Usage() {
	fmt.Fprintf(os.Stderr, `
Usage: go run cmd/migrate/main.go -c {配置文件} -a {up(迁移)/down(回滚)}
`)
}
