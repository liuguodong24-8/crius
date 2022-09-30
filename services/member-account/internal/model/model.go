package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/redis"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/member-account/internal/config"
	"gorm.io/gorm"
)

var entity *db.Entity
var redisEntity *redis.Entity

// SnapShotChan 快照channel
var SnapShotChan chan *TableSnapshot

// DatabaseConnection 数据库连接
func DatabaseConnection() error {
	cfg := config.Setting.Database
	var err error

	entity, err = db.NewEntity(db.Config{
		Dialect:  "postgres",
		Server:   cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Database: cfg.DBName,
		Password: cfg.Password,
		Debug:    cfg.Log,
	})

	SnapShotChan = make(chan *TableSnapshot, 200)
	go snapShotRec()
	return err
}

func GetDBEntity() *db.Entity {
	return entity
}

// Payments 支付方式
type Payments struct {
	Wechat int32 `json:"wechat"`
	Cash   int32 `json:"cash"`
	Card   int32 `json:"card"`
	Alipay int32 `json:"alipay"`
}

//Value 转换json存储
func (p Payments) Value() (driver.Value, error) {
	return json.Marshal(p)
}

//Scan 将数据库对象转换成可以使用的golang 属性
func (p *Payments) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("payments error")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("payments json error")
	}

	return nil
}

// DatabaseConn 返回外部调用conn
func DatabaseConn() *db.Entity {
	return entity
}

func snapShotRec() {
	for snapShot := range SnapShotChan {
		if err := CreateSnapshot(*snapShot); err != nil {
			cutil.Logger.Error(fmt.Sprintf("创建快照错误:%v", err))
		}
	}
}

// RedisConnection redis连接
func RedisConnection() (err error) {
	redisEntity, err = redis.NewEntity(redis.Config{
		IP:       config.Setting.Redis.IP,
		Port:     config.Setting.Redis.Port,
		Password: config.Setting.Redis.Password,
		Database: config.Setting.Redis.Database,
	})
	return err
}

func idCondition(id uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func pagingCondition(offset, limit int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			return db
		}
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

func statusCondition(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" {
			db = db.Where("status = ?", status)
		}
		return db
	}
}

func merchantIDCondition(id uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if id != uuid.Nil {
			db = db.Where("merchant_id = ?", id)
		}
		return db
	}
}
