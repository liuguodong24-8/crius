package model

import (
	"fmt"
	"time"

	"gorm.io/gorm/clause"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/config"
	"gorm.io/gorm"
)

var entity *db.Entity

// SnapShotChan 快照channel
var SnapShotChan chan *TableSnapshot

const (
	// StatusOpened 状态opened
	StatusOpened = "opened"
	// StatusClosed 状态closed
	StatusClosed = "closed"
)

// DatabaseConnection 数据库连接
func DatabaseConnection() error {
	cfg := config.Setting.DataBase
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

	//插件注册
	entity.Conn.Callback().Query().Before("gorm:query").Register("phone_query", phoneQuery)

	SnapShotChan = make(chan *TableSnapshot, 200)
	go snapShotRec()
	return err
}

// DatabaseConn 外部获取调用
func DatabaseConn() *gorm.DB {
	return entity.Conn
}

func snapShotRec() {
	for snapShot := range SnapShotChan {
		if err := CreateSnapshot(*snapShot); err != nil {
			crius.Logger.Error(fmt.Sprintf("创建快照错误:%v", err))
		}
	}
}

// GetStaffSequence 获取员工编号自增序列
func GetStaffSequence() int64 {
	m := make(map[string]interface{})
	if err := entity.Conn.Raw("select nextval('merchant_basic.staff_code_seq')").Scan(&m).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetStaffSequence 获取员工sequence数据库错误:%v", err))
		return -1
	}
	seq, ok := m["nextval"].(int64)
	if !ok {
		crius.Logger.Error("GetStaffSequence 获取员工sequence数据库失败")
	}
	return seq
}

// GetBranchSequence 获取门店编号自增序列
func GetBranchSequence() int64 {
	m := make(map[string]interface{})
	if err := entity.Conn.Raw("select nextval('merchant_basic.branch_code_seq')").Scan(&m).Error; err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchSequence 获取门店sequence数据库错误:%v", err))
		return -1
	}
	seq, ok := m["nextval"].(int64)
	if !ok {
		crius.Logger.Error("GetBranchSequence 获取门店sequence数据库失败")
	}
	return seq
}

func pagingCondition(offset, limit int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			return db
		}
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

func getTodayDate(second int64) time.Time {
	t := time.Unix(second, 0)
	return t.Add(-time.Hour * time.Duration(t.Hour())).Add(-time.Minute * time.Duration(t.Minute())).Add(-time.Second * time.Duration(t.Second())).Add(-time.Nanosecond * time.Duration(t.Nanosecond()))
}

func getTomorrowDate(second int64) time.Time {
	t := time.Unix(second, 0)
	return t.Add(time.Hour * 24).Add(-time.Hour * time.Duration(t.Hour())).Add(-time.Minute * time.Duration(t.Minute())).Add(-time.Second * time.Duration(t.Second())).Add(-time.Nanosecond * time.Duration(t.Nanosecond()))
}

func phoneQuery(db *gorm.DB) {
	schema := db.Statement.Schema
	if schema != nil && schema.Name == "TableMember" {
		if c, ok := db.Statement.Clauses["WHERE"]; ok {
			if where, ok := c.Expression.(clause.Where); ok {
				for _, expr := range where.Exprs {
					if orConf, ok := expr.(clause.Expr); ok {
						sql := orConf.SQL
						if (len(sql) > 7) && (sql[:7] == "phone =") {
							phone := orConf.Vars[0].(string)
							phoneTail := phone[len(phone)-1:]
							var vars []interface{}
							vars = append(vars, phoneTail)
							clauseTail := clause.Expr{
								SQL:                "phone_tail = ?",
								Vars:               vars,
								WithoutParentheses: false,
							}
							where.Exprs = append(where.Exprs, clauseTail)
							c.Expression = where
							db.Statement.Clauses["WHERE"] = c
						}
					}
				}
			}
		}
	}
}
