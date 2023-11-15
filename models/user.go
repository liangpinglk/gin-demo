package models

import (
	"database/sql"
	"time"
)

type UserInfo struct {
	Birthday string `json:"birthday"`
	Job      string `json:"job"`
	Sex      string `json:"sex"`
}

type CommonInfo struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CommonInfo   CommonInfo `gorm:"embedded"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user_info"
}
