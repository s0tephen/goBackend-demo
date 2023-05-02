// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePost = "post"

// Post mapped from table <post>
type Post struct {
	PID      int32     `gorm:"column:pId;primaryKey;autoIncrement:true" json:"pId"`
	Username *string   `gorm:"column:username" json:"username"`
	PTitle   string    `gorm:"column:pTitle;not null" json:"pTitle"`
	PCenter  string    `gorm:"column:pCenter;not null" json:"pCenter"`
	PImg     string    `gorm:"column:pImg;not null" json:"pImg"`
	PLabel   string    `gorm:"column:pLabel;not null" json:"pLabel"`
	PTime    time.Time `gorm:"column:pTime;not null" json:"pTime"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}
