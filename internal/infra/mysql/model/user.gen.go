// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户ID" json:"id"`                       // 用户ID
	Username  string    `gorm:"column:username;not null;uniqueIndex:username,priority:1;comment:用户名" json:"username"` // 用户名
	Password  string    `gorm:"column:password;not null;comment:密码（建议使用哈希存储）" json:"password"`                        // 密码（建议使用哈希存储）
	Avatar    string    `gorm:"column:avatar;not null;comment:头像URL" json:"avatar"`                                   // 头像URL
	OpenID    string    `gorm:"column:open_id;not null;comment:微信openId" json:"open_id"`                              // 微信openId
	Phone     string    `gorm:"column:phone;not null" json:"phone"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
