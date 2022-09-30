package model

import uuid "github.com/satori/go.uuid"

// TableFileTransfer 测试文件写入
type TableFileTransfer struct {
	ID       uuid.UUID `gorm:"column:id"`
	Filename string    `json:"filename"`
	Contents []byte    `json:"contents"`
}

// TableName 指定表名
func (TableFileTransfer) TableName() string {
	return "merchant_basic.file_transfer"
}
