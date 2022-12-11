// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameLink = "link"

// Link mapped from table <link>
type Link struct {
	ID          int32     `gorm:"column:id;type:integer;primaryKey" json:"id"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Logo        string    `gorm:"column:logo;type:varchar(1023);not null" json:"logo"`
	Name        string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Priority    int32     `gorm:"column:priority;type:integer;not null" json:"priority"`
	Team        string    `gorm:"column:team;type:varchar(255);not null" json:"team"`
	URL         string    `gorm:"column:url;type:varchar(1023);not null" json:"url"`
}

// TableName Link's table name
func (*Link) TableName() string {
	return TableNameLink
}
