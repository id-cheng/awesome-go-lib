package model

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint // Standard field for the primary key
	UUID         string
	Name         string // 一个常规字符串字段
	Role         string
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          `gorm:"default:18"`   // 默认值18
	Active       sql.NullBool   `gorm:"default:true"` // 默认值true
	Birthday     time.Time      // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}

// BeforeCreate 创建钩子 通过实现这些接口 BeforeSave, BeforeCreate, AfterSave, AfterCreate来自定义钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()
	if u.Role == "admin" {
		return errors.New("invalid role")
	}
	return
}
