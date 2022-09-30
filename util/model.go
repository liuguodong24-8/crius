package util

import (
	"fmt"
	"reflect"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Status 启用禁用状态
type Status string

const (
	// StatusOpened 启用
	StatusOpened Status = `opened`
	// StatusClosed 禁用
	StatusClosed Status = `closed`
)

// String ...
func (s Status) String() string {
	return string(s)
}

// StringToStatus 字符串转type
func StringToStatus(s string) Status {
	return Status(s)
}

// ColumnNotEqualScope 字段不等于
func ColumnNotEqualScope(column string, search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		return db.Where(fmt.Sprintf("%s != ?", column), search)
	}
}

// ColumnSymbolScope 字段条件
func ColumnSymbolScope(column string, symbol string, search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s %s ?", column, symbol), search)
	}
}

// ColumnEqualScope 字段等于
func ColumnEqualScope(column string, search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", column), search)
	}
}

// ColumnEqualScopeDefault scope条件如果value为默认值，不筛选该条件
func ColumnEqualScopeDefault(column string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		valueEmpty := true
		switch value.(type) {
		case uuid.UUID:
			if value.(uuid.UUID) != uuid.Nil {
				valueEmpty = false
			}
		case string:
			if value.(string) != "" {
				valueEmpty = false
			}
		case int:
			if value.(int) != 0 {
				valueEmpty = false
			}
		case int8:
			if value.(int8) != 0 {
				valueEmpty = false
			}
		case int16:
			if value.(int16) != 0 {
				valueEmpty = false
			}
		case int32:
			if value.(int32) != 0 {
				valueEmpty = false
			}
		case int64:
			if value.(int64) != 0 {
				valueEmpty = false
			}
		case uint8:
			if value.(uint8) != 0 {
				valueEmpty = false
			}
		case uint16:
			if value.(uint16) != 0 {
				valueEmpty = false
			}
		case uint32:
			if value.(uint32) != 0 {
				valueEmpty = false
			}
		case uint64:
			if value.(uint64) != 0 {
				valueEmpty = false
			}
		}
		if !valueEmpty {
			db = db.Where(fmt.Sprintf("%s = ?", column), value)
		}
		return db
	}
}

// ColumnNullScope 是否为null查找
func ColumnNullScope(column string, isNull bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if isNull {
			return db.Where(fmt.Sprintf(`%s is null`, column))
		}

		return db.Where(fmt.Sprintf(`%s is not null`, column))
	}
}

// ColumnInScope 字段in
func ColumnInScope(column string, values []interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s IN ?", column), values)
	}
}

// ColumnInScopeDefault scope条件如果value数组等于0，不筛选该条件
func ColumnInScopeDefault(column string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value != nil {
			t := reflect.TypeOf(value)
			v := reflect.ValueOf(value)
			if (t.Kind() == reflect.Array || t.Kind() == reflect.Slice) && v.Len() > 0 {
				db = db.Where(fmt.Sprintf("%s in (?)", column), value)
			}
		}
		return db
	}
}

// ColumnLikeScope 模糊搜索
func ColumnLikeScope(column, search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(column) == 0 || len(search) == 0 {
			return db
		}

		return db.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", search))
	}
}

// PaginationScope 分页
func PaginationScope(offset, limit int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			limit = 20
		}

		if offset <= 0 {
			offset = 0
		}
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

// ArrayAnyScope 查询数组某个值
func ArrayAnyScope(column string, search interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("? = ANY(%s)", column), search)
	}
}

// ArrayOverlapScope 查询数组 && 数组
func ArrayOverlapScope(column string, t string, values []interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var str []string
		for _, val := range values {
			str = append(str, fmt.Sprintf("'%v'", val))
		}

		return db.Where(fmt.Sprintf("%s && ARRAY[?]::%s[]", column, t), gorm.Expr(strings.Join(str, `,`)))
	}
}
