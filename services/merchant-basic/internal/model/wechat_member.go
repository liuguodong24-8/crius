package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
)

//TableWechatMember 微信用户表
type TableWechatMember struct {
	ID         uuid.UUID    `gorm:"column:id"`
	OpenID     string       `gorm:"column:openid"`
	AppID      string       `gorm:"column:appid"`
	MemberID   *uuid.UUID   `gorm:"column:member_id"`
	Nickname   string       `gorm:"column:nickname"`
	Sex        int8         `gorm:"column:sex"`
	Province   string       `gorm:"column:province"`
	City       string       `gorm:"column:city"`
	HeadImgURL string       `gorm:"column:headimgurl"`
	UnionID    string       `gorm:"column:unionid"`
	Extra      *pkgs.Params `gorm:"column:extra"`
	CreatedAt  *time.Time   `gorm:"column:created_at"`
	UpdatedAt  *time.Time   `gorm:"column:updated_at"`
	Member     TableMember  `gorm:"foreignkey:MemberID"`
	LoadExtra  string       `gorm:"-"` // 导入信息 不处理
}

//TableName 微信用户表
func (t *TableWechatMember) TableName() string {
	return "merchant_basic.wechat_member"
}

// ShowWechatUser ...
func ShowWechatUser(ID uuid.UUID) (*TableWechatMember, error) {
	var user TableWechatMember
	err := entity.Conn.Preload("Member").
		Scopes(
			crius.ColumnEqualScope("id", ID),
		).First(&user).Error
	return &user, err
}

//ShowWechatUserByMember 查询微信用户信息
func ShowWechatUserByMember(memberID uuid.UUID) (*TableWechatMember, error) {
	var user TableWechatMember
	err := entity.Conn.
		Preload("Member").
		Scopes(
			crius.ColumnEqualScope("member_id", memberID),
		).First(&user).Error
	return &user, err
}

//ShowWechatUserByOpenID 查询
func ShowWechatUserByOpenID(OpenID, AppID string) (*TableWechatMember, error) {
	var user TableWechatMember
	err := entity.Conn.
		Preload("Member").
		Scopes(
			crius.ColumnEqualScope("openid", OpenID),
			crius.ColumnEqualScope("appid", AppID),
		).First(&user).Error
	return &user, err
}

//ShowWechatUserByUnionID 查询
func ShowWechatUserByUnionID(UnionID, AppID string) (*TableWechatMember, error) {
	var user TableWechatMember
	err := entity.Conn.
		Preload("Member").
		Scopes(
			crius.ColumnEqualScope("appid", AppID),
			crius.ColumnEqualScope("unionid", UnionID),
		).First(&user).Error
	return &user, err
}

//UpdateWechatUser 更新
func UpdateWechatUser(AppID, UnionID string, wechatUser *TableWechatMember) error {
	return entity.Conn.Model(&TableWechatMember{}).Scopes(
		crius.ColumnEqualScope("appid", AppID),
		crius.ColumnEqualScope("unionid", UnionID),
	).Updates(wechatUser).Error
}

//CreateWechatUser 创建
func CreateWechatUser(wechatUser *TableWechatMember) error {
	return entity.Conn.Create(wechatUser).Error
}

// GetWechatUsers list
func GetWechatUsers(ids []uuid.UUID) ([]TableWechatMember, error) {
	var users []TableWechatMember
	err := entity.Conn.Select("id, nickname, headimgurl, member_id").Where("id in (?)", ids).Find(&users).Error
	return users, err
}
