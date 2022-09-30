package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gorm.io/gorm"
)

// TableAppointment 预约表结构
type TableAppointment struct {
	ID               uuid.UUID               `gorm:"column:id"`
	MerchantID       *uuid.UUID              `gorm:"column:merchant_id"`
	BranchID         *uuid.UUID              `gorm:"column:branch_id"`
	CalledCode       string                  `gorm:"column:called_code"`
	CalledPhone      string                  `gorm:"column:called_phone"`
	AppointmentCode  string                  `gorm:"column:appointment_code"`
	AppointmentPhone string                  `gorm:"column:appointment_phone"`
	Name             string                  `gorm:"column:name"`
	Gender           int8                    `gorm:"column:gender"`
	Way              int8                    `gorm:"column:way"`
	CustomerNum      int16                   `gorm:"column:customer_num"`
	AppointmentAt    *time.Time              `gorm:"column:appointment_at"`
	AppointmentDate  *time.Time              `gorm:"column:appointment_date"`
	ExpireAt         *time.Time              `gorm:"column:expire_at"`
	RoomTypeID       *uuid.UUID              `gorm:"column:room_type_id"`
	DepositFee       int32                   `gorm:"column:deposit_fee"`
	FlowerCake       bool                    `gorm:"column:flower_cake"`
	FlowerCakeRemark string                  `gorm:"column:flower_cake_remark"`
	Remark           string                  `gorm:"column:remark"`
	Operator         *uuid.UUID              `gorm:"column:operator"`
	Status           AppointmentStatus       `gorm:"column:status"`
	CancelledReason  string                  `gorm:"column:cancelled_reason"`
	CreatedAt        *time.Time              `gorm:"column:created_at"`
	UpdatedAt        *time.Time              `gorm:"column:updated_at"`
	RelatedID        *uuid.UUID              `gorm:"column:related_id"`
	DateCounterID    *uuid.UUID              `gorm:"column:date_counter_id"`
	ChargingWay      int8                    `gorm:"column:charging_way"`
	TradeID          *uuid.UUID              `gorm:"column:trade_id"`
	KeepAt           *time.Time              `gorm:"column:keep_at"`
	CancelAt         *time.Time              `gorm:"column:cancel_at"`
	TradeType        string                  `gorm:"column:trade_type"`
	ArrivedAt        *time.Time              `gorm:"column:arrived_at"`
	DeletedAt        gorm.DeletedAt          `gorm:"column:deleted_at"`
	MemberID         *uuid.UUID              `gorm:"column:member_id"`
	Breach           bool                    `gorm:"column:breach"`
	BreachReason     AppointmentBreachReason `gorm:"column:breach_reason"`
	WechatIDs        *WechatIDs              `gorm:"column:wechat_ids"`
}

// WechatID 微信id
type WechatID struct {
	ID   uuid.UUID `json:"id"`
	Time time.Time `json:"time"`
}

// WechatIDs 微信id数组
type WechatIDs []WechatID

// Value 返回数据库可识别类型
func (p WechatIDs) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// Scan ...
func (p *WechatIDs) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("util params value")
	}

	if err := json.Unmarshal(bytes, p); nil != err {
		return errors.New("util params value error")
	}

	return nil
}

// Appointment 关联退款预约结构
type Appointment struct {
	TableAppointment
	TableAppointmentExtend
}

type (
	// AppointmentStatus 预约状态
	AppointmentStatus string
	// AppointmentBreachReason 违约原因
	AppointmentBreachReason string
)

const (
	//AppointmentStatusArrearage 待支付
	AppointmentStatusArrearage AppointmentStatus = "arrearage"
	//AppointmentStatusAppointed 已预约
	AppointmentStatusAppointed AppointmentStatus = "appointed"
	//AppointmentStatusArrived 已到店
	AppointmentStatusArrived AppointmentStatus = "arrived"
	//AppointmentStatusCancelled 已取消
	AppointmentStatusCancelled AppointmentStatus = "cancelled"
	//AppointmentStatusRefunding 退款中
	AppointmentStatusRefunding AppointmentStatus = "refunding"
	//AppointmentStatusRefunded 退款成功
	AppointmentStatusRefunded AppointmentStatus = "refunded"
	//AppointmentStatusExpired 已过期
	AppointmentStatusExpired AppointmentStatus = "expired"

	//AppointmentChangeCancelReason 预约改变记录默认取消原因
	AppointmentChangeCancelReason = "系统操作"

	// AppointmentBreachReasonLate 迟到
	AppointmentBreachReasonLate AppointmentBreachReason = "迟到"
	// AppointmentBreachReasonNotArrive 未到店
	AppointmentBreachReasonNotArrive AppointmentBreachReason = "未到店"
	// AppointmentBreachReasonCancelLate 过晚取消
	AppointmentBreachReasonCancelLate AppointmentBreachReason = "过晚取消"
)

// BranchAppointmentBusinessTimeKey 营业时间缓存
const BranchAppointmentBusinessTimeKey = "branch:appointment:business_time:branch_id:%s:date:%d"

// TableName 表名
func (TableAppointment) TableName() string {
	return "appointment.appointment"
}

// UpdateAppointmentWechatIDs 更新预约分享微信id
func UpdateAppointmentWechatIDs(a *TableAppointment) error {
	return entity.Conn.Model(&TableAppointment{}).Scopes(util.ColumnEqualScope("id", a.ID)).UpdateColumn("wechat_ids", a.WechatIDs).Error
}

// ShowAppointment 查询预约
func ShowAppointment(id uuid.UUID) (*TableAppointment, error) {
	appointment := new(TableAppointment)
	err := entity.Conn.Model(&TableAppointment{}).Scopes(util.ColumnEqualScope("id", id)).Take(appointment).Error
	return appointment, err
}

// ShowFullAppointment 查询完整预约
func ShowFullAppointment(id uuid.UUID) (*Appointment, error) {
	appointment := new(Appointment)
	err := entity.Conn.Model(&TableAppointment{}).Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").Scopes(util.ColumnEqualScope("appointment.id", id)).Take(appointment).Error
	return appointment, err
}

// OpenRoom 开房
func OpenRoom(id, roomID uuid.UUID, openAt time.Time) error {
	tx := entity.Conn.Begin()
	m := make(map[string]interface{})
	m["open_room_id"] = roomID
	m["open_at"] = openAt
	if err := tx.Model(&TableAppointmentExtend{}).Scopes(util.ColumnEqualScope("appointment_id", id)).Updates(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&TableAppointment{}).Scopes(util.ColumnEqualScope("id", id)).UpdateColumn("status", string(AppointmentStatusArrived)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// CancelOpenRoom 取消开房
func CancelOpenRoom(appointment *Appointment) error {
	tx := entity.Conn.Begin()
	if err := tx.Scopes(util.ColumnEqualScope("id", appointment.ID)).Select("arrived_at", "status", "breach", "breach_reason").Updates(&(appointment.TableAppointment)).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Scopes(util.ColumnEqualScope("appointment_id", appointment.ID)).Select("open_at", "open_room_id").Updates((appointment.TableAppointmentExtend)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// GetAppointmentsBySended 根据是否发送http请求获取预约列表
func GetAppointmentsBySended(t time.Time, after, sended bool) ([]Appointment, error) {
	appointments := make([]Appointment, 0)
	db := entity.Conn.Model(&TableAppointment{}).Joins("left join appointment.appointment_extend on appointment.id = appointment_extend.appointment_id").
		Scopes(util.ColumnEqualScope("appointment_extend.sended", sended), util.ColumnInScope("appointment.status", []interface{}{AppointmentStatusAppointed, AppointmentStatusArrived, AppointmentStatusCancelled, AppointmentStatusRefunded, AppointmentStatusRefunding}))
	if after {
		db = db.Scopes(util.LargeScope("appointment.appointment_at", t))
	} else {
		db = db.Scopes(DateRangeCondition(t, t.Add(time.Hour*24)))
	}
	err := db.Find(&appointments).Error
	return appointments, err
}

// AppointmentPhoneTailLike 手机号后缀模糊查询
func AppointmentPhoneTailLike(suf string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if suf != "" {
			return db.Where("appointment.appointment_phone like ?", fmt.Sprintf("%%%s", suf))
		}
		return db
	}
}

// WayAnd 预约方式与
func WayAnd(way int8) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if way != 0 {
			return db.Where("way & ? = ?", way, way)
		}
		return db
	}
}

// IsDepositPayed 预约是否支付订金
func IsDepositPayed(payed int8) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if payed == 1 {
			return db.Where("deposit_fee != 0")
		} else if payed == 2 {
			return db.Where("deposit_fee = 0")
		}
		return db
	}
}

// FlowerCakeCondition 预约是否代收鲜花蛋糕
func FlowerCakeCondition(flowerCake int8) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if flowerCake == 1 {
			return db.Where("appointment.flower_cake = ?", true)
		} else if flowerCake == 2 {
			return db.Where("appointment.flower_cake = ?", false)
		}
		return db
	}
}

// DateRangeCondition 预约时间
func DateRangeCondition(dateStart, dateEnd time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !dateStart.IsZero() && !dateEnd.IsZero() && !dateEnd.Before(dateStart) {
			return db.Where("appointment.appointment_at >= ? and appointment.appointment_at <= ?", dateStart, dateEnd)
		}
		return db
	}
}

// PagingCondition 分页
func PagingCondition(offset, limit int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit > 0 {
			db = db.Offset(int(offset)).Limit(int(limit))
		}
		return db
	}
}

// AnyPhoneMemberIDCondition 查询任一电话或者会员id
func AnyPhoneMemberIDCondition(phone string, memberID uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var sql string
		if len(phone) != 0 {
			sql += fmt.Sprintf("appointment.appointment_phone = '%s' or appointment.called_phone = '%s'", phone, phone)
		}
		if memberID != uuid.Nil {
			if len(sql) != 0 {
				sql += " or "
			}
			sql += fmt.Sprintf(`appointment.member_id = '%s'`, memberID.String())
		}
		if len(sql) != 0 {
			db = db.Where(sql)
		}
		return db
	}
}

// AnyPhoneMemberIDWechatIDCondition 查询任一电话或者会员id或微信id
func AnyPhoneMemberIDWechatIDCondition(phone string, memberID, wechatID uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var sql string
		if len(phone) != 0 {
			sql += fmt.Sprintf("appointment.appointment_phone = '%s' or appointment.called_phone = '%s'", phone, phone)
		}
		if memberID != uuid.Nil {
			if len(sql) != 0 {
				sql += " or "
			}
			sql += fmt.Sprintf(`appointment.member_id = '%s'`, memberID.String())
		}
		if wechatID != uuid.Nil {
			if len(sql) != 0 {
				sql += " or "
			}
			sql += fmt.Sprintf(`appointment.wechat_ids @> '[{"id":"%s"}]'`, wechatID.String())
		}
		if len(sql) != 0 {
			db = db.Where(sql)
		}
		return db
	}
}

// ConfigureTimeCondition 预约模板时间
func ConfigureTimeCondition(t fields.LocalTime, isNextDay bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if isNextDay {
			db = db.Where("(c::json->>'time')::time >= ? and (c::json->>'is_next_day')::bool = ?", t, isNextDay)
		} else {
			db = db.Where("((c::json->>'time')::time >= ? and (c::json->>'is_next_day')::bool = ?) or ((c::json->>'time')::time < ? and (c::json->>'is_next_day')::bool = ?)", t, isNextDay, t, true)
		}
		return db
	}
}

// ConfigureWayAnd 预约模板预约方式与
func ConfigureWayAnd(way int8) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("(c::json->>'way')::int & ? = ?", way, way)
	}
}
