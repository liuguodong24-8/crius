package model

import (
	"bytes"
	"database/sql/driver"
	"strings"
	"time"

	crius "gitlab.omytech.com.cn/micro-service/Crius/util"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// TableMember member表
type TableMember struct {
	ID            uuid.UUID         `gorm:"column:id"`
	Name          string            `gorm:"column:name"`
	Phone         string            `gorm:"column:phone"`
	PhoneTail     string            `gorm:"column:phone_tail"`
	PhoneSuffix   string            `gorm:"column:phone_suffix"`
	PhoneCode     string            `gorm:"column:phone_code"`
	Gender        int8              `gorm:"column:gender"`
	City          string            `gorm:"column:city"`
	Avatar        string            `gorm:"column:avatar"`
	Birthday      *time.Time        `gorm:"column:birthday"`
	Code          string            `gorm:"column:code"`
	FirstBranchID *uuid.UUID        `gorm:"column:first_branch_id"`
	FirstBrand    *uuid.UUID        `gorm:"column:first_brand"`
	StaffID       *uuid.UUID        `gorm:"column:staff_id"`
	FirstChannel  MemberChannel     `gorm:"column:first_channel"`
	Channels      *MemberChannelArr `gorm:"column:channels"`
	MerchantID    *uuid.UUID        `gorm:"column:merchant_id"`
	DeletedAt     gorm.DeletedAt    `gorm:"column:deleted_at"`
	CreatedAt     *time.Time        `gorm:"column:created_at"`
	UpdatedAt     *time.Time        `gorm:"column:updated_at"`
	LoadExtra     string            `gorm:"-"` // 导入信息 不处理
}

// TableName 指定表名
func (TableMember) TableName() string {
	return "merchant_basic.member"
}

// MemberChannel ...
type MemberChannel string

type MemberChannelArr []MemberChannel

// Value 转换为数据库字段类型
func (cha MemberChannelArr) Value() (driver.Value, error) {
	var buffer bytes.Buffer

	buffer.WriteString("{")
	for k, s := range cha {
		if k > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(s))
	}
	buffer.WriteString("}")

	return buffer.String(), nil
}

// Slice slice
func (cha *MemberChannelArr) Slice() []string {
	if cha != nil {
		var response []string
		for _, s := range *cha {
			response = append(response, string(s))
		}

		return response
	}
	return nil
}

// Scan 将数据库数据映射到结构体
func (cha *MemberChannelArr) Scan(src interface{}) error {
	if nil == src {
		return nil
	}

	val := src.(string)
	value := strings.Trim(strings.Trim(val, "{"), "}")
	if len(value) == 0 {
		return nil
	}

	var res MemberChannelArr
	for _, v := range strings.Split(value, ",") {
		if len(v) == 0 {
			continue
		}
		res = append(res, MemberChannel(v))
	}

	*cha = res

	return nil
}

const (
	//MemberChannelOpenCard 开卡
	MemberChannelOpenCard MemberChannel = "open_card"

	//MemberChannelWechat 微信
	MemberChannelWechat MemberChannel = "wechat"

	//MemberChannelCall 来电
	MemberChannelCall MemberChannel = "call"

	//MemberChannelCallAppointment 来电预约
	MemberChannelCallAppointment MemberChannel = "call_appointment"

	//MemberChannelStaffAppointment 客服预约
	MemberChannelStaffAppointment MemberChannel = "staff_appointment"
)

//MemberWithBehavior 包含交互次数
type MemberWithBehavior struct {
	TableMember
	BehaviorCount int32 `gorm:"column:behavior_count"`
}

// BeforeCreate member新建hook
func (m *TableMember) BeforeCreate(tx *gorm.DB) (err error) {
	m.PhoneTail = m.Phone[len(m.Phone)-1:]
	m.PhoneSuffix = m.Phone[len(m.Phone)-4:]
	return nil
}

// BeforeSave member保存hook
func (m *TableMember) BeforeSave(tx *gorm.DB) (err error) {
	if m.Phone != "" {
		m.PhoneTail = m.Phone[len(m.Phone)-1:]
		m.PhoneSuffix = m.Phone[len(m.Phone)-4:]
	}
	return nil
}

// BeforeUpdate member更新hook
func (m *TableMember) BeforeUpdate(tx *gorm.DB) (err error) {
	if m.Phone != "" {
		m.PhoneTail = m.Phone[len(m.Phone)-1:]
		m.PhoneSuffix = m.Phone[len(m.Phone)-4:]
	}
	return nil
}

// CreateMember 新增会员
func CreateMember(member TableMember) error {
	return entity.Conn.Omit("code").Create(&member).Error
}

// UpdateMember 更新会员信息
func UpdateMember(member *TableMember) error {
	return entity.Conn.Updates(member).Error
}

// GetMembers 会员列表
func GetMembers(scopes []func(*gorm.DB) *gorm.DB) ([]MemberWithBehavior, error) {
	var members []MemberWithBehavior
	sub := entity.Conn.Model(&TableMemberBehavior{}).Select("count(*) mb_count, member_id").Group("member_id")
	err := entity.Conn.
		Select("member.*, mb.mb_count as behavior_count").
		Joins("left join (?) mb on mb.member_id = member.id", sub).
		Scopes(scopes...).Order("created_at desc").
		Find(&members).Error

	return members, err
}

// CountMembers 计数
func CountMembers(scopes []func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := entity.Conn.Model(&TableMember{}).Scopes(scopes...).Count(&count).Error

	return count, err
}

// GetMembersByIDs 批量获取会员信息
func GetMembersByIDs(ids []uuid.UUID) ([]TableMember, error) {
	var members []TableMember
	err := entity.Conn.Scopes(
		crius.ColumnInScopeDefault("id", ids),
	).Order("created_at desc").
		Find(&members).Error

	return members, err
}

// GetMembersByPhoneSuffix 根据手机尾号查询会员
func GetMembersByPhoneSuffix(phoneSuffix string, merchantID uuid.UUID) ([]TableMember, error) {
	var members []TableMember
	err := entity.Conn.Scopes(
		crius.ColumnEqualScope("phone_suffix", phoneSuffix),
		crius.ColumnEqualScope("merchant_id", merchantID),
	).
		Order("created_at desc").
		Find(&members).Error

	return members, err
}

// ShowMember 获取单个会员信息
func ShowMember(id uuid.UUID) (*TableMember, error) {
	var member TableMember
	err := entity.Conn.Where("id = ?", id).Take(&member).Error
	return &member, err
}

// ShowMemberByAccuratePhone 手机号精确查找用户
func ShowMemberByAccuratePhone(phoneCode, phone string, merchantID uuid.UUID) (*TableMember, error) {
	var member TableMember
	if phone == "" {
		return nil, gorm.ErrRecordNotFound
	}
	err := entity.Conn.Unscoped().Scopes(
		memberPhoneCondition(phone, phoneCode),
		crius.ColumnEqualScope("merchant_id", merchantID),
	).Take(&member).Error
	return &member, err
}

// ShowMemberExistsByPhone 会员是否存在
func ShowMemberExistsByPhone(phone, phoneCode string, merchantID uuid.UUID) error {
	if phone == "" {
		return gorm.ErrRecordNotFound
	}
	return entity.Conn.Unscoped().Scopes(
		memberPhoneCondition(phone, phoneCode),
		crius.ColumnEqualScope("merchant_id", merchantID),
	).Take(&TableMember{}).Error
}

// GetBirthdayMembers 获取生日的会员
func GetBirthdayMembers(birthday string, merchantID uuid.UUID) ([]TableMember, error) {
	var members []TableMember
	err := entity.Conn.Scopes(crius.ColumnEqualScope("merchant_id", merchantID), birthdayCondition(birthday)).Find(&members).Error
	return members, err
}

func memberPhoneCondition(phone, phoneCode string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if phone != "" {
			db = db.Where("phone = ?", phone)
		}
		//不再校验手机区号，不影响唯一性
		return db
	}
}

func birthdayCondition(birthday string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(gorm.Expr("to_char(birthday, 'MM-dd') = ?", birthday))
	}
}

// GetCouponMemberIDs 推券对象
func GetCouponMemberIDs(scopes []func(db *gorm.DB) *gorm.DB) ([]TableMember, error) {
	var members []TableMember
	err := entity.Conn.Model(TableMember{}).Select("id, created_at").Scopes(scopes...).Order("created_at desc").Find(&members).Error
	return members, err
}

// CountCouponMemberIDs count
func CountCouponMemberIDs(scopes []func(db *gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := entity.Conn.Model(TableMember{}).Scopes(scopes...).Count(&count).Error
	return count, err
}

// SearchMember 会员查询
func SearchMember(scopes []func(*gorm.DB) *gorm.DB) ([]TableMember, error) {
	var members []TableMember
	err := entity.Conn.Scopes(scopes...).Order("created_at desc").Limit(50).Find(&members).Error
	return members, err
}
