// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameCommentBlack = "comment_black"

// CommentBlack mapped from table <comment_black>
type CommentBlack struct {
	ID         int32     `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	BanTime    time.Time `gorm:"column:ban_time;type:datetime;not null" json:"ban_time"`
	IPAddress  string    `gorm:"column:ip_address;type:varchar(127);not null" json:"ip_address"`
}

// TableName CommentBlack's table name
func (*CommentBlack) TableName() string {
	return TableNameCommentBlack
}
