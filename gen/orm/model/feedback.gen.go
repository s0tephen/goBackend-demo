// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFeedback = "feedback"

// Feedback mapped from table <feedback>
type Feedback struct {
	Fid   int32     `gorm:"column:fid;primaryKey;autoIncrement:true" json:"fid"`          // 反馈ID
	FUser string    `gorm:"column:fUser;not null" json:"fUser"`                           // 反馈者
	FMsg  string    `gorm:"column:fMsg;not null" json:"fMsg"`                             // 反馈内容
	FTime time.Time `gorm:"column:fTime;not null;default:CURRENT_TIMESTAMP" json:"fTime"` // 反馈时间
}

// TableName Feedback's table name
func (*Feedback) TableName() string {
	return TableNameFeedback
}