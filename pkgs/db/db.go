package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ==========================
// 参考  https://github.com/go-gorm/playground/blob/master/db.go
// 后续考虑 在 interface 定义需要实现的方法
// Conn 私有化 接口注册调用
// =======================

// Interface interface
type Interface interface {
	// todo
}

// Dialect db driver
type Dialect string

// DialectPostgre postgre
const DialectPostgre Dialect = "postgres"

// DialectMysql mysql
const DialectMysql Dialect = "mysql"

// DialectSqlite sqlite
const DialectSqlite Dialect = "sqlite"

// DialectSqlserver sql server
const DialectSqlserver Dialect = "sqlserver"

// Config 数据库配置
type Config struct {
	Dialect  Dialect
	Server   string
	Port     int
	User     string
	Database string
	Password string
	Debug    bool
}

// Entity 对外暴露实例
type Entity struct {
	Conn *gorm.DB
}

// NewEntity 新建实例
func NewEntity(cfg Config) (*Entity, error) {
	db, err := build(cfg)
	if err != nil {
		return nil, err
	}

	return &Entity{Conn: db}, nil
}

func build(cfg Config) (*gorm.DB, error) {
	gormConfig := &gorm.Config{}
	if cfg.Debug {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	if cfg.Dialect == DialectPostgre {
		return gorm.Open(postgres.Open(pgsqlDsn(cfg)), gormConfig)
	}

	if cfg.Dialect == DialectMysql {
		return gorm.Open(mysql.Open(mysqlDsn(cfg)), gormConfig)
	}

	if cfg.Dialect == DialectSqlite {
		return gorm.Open(sqlite.Open(sqliteDsn(cfg)), gormConfig)
	}

	if cfg.Dialect == DialectSqlserver {
		return gorm.Open(sqlserver.Open(sqlServerDsn(cfg)), gormConfig)
	}

	return nil, fmt.Errorf("无法匹配数据库连接驱动, :%v", cfg.Dialect)
}

func pgsqlDsn(cfg Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Chongqing",
		cfg.Server, cfg.User, cfg.Password, cfg.Database, cfg.Port)
}

func mysqlDsn(cfg Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.Database)
}

func sqliteDsn(config Config) string {
	return config.Server
}

func sqlServerDsn(config Config) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s;encrypt=disable;",
		config.User, config.Password, config.Server, config.Port, config.Database)
}
