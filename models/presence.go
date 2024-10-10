package models

import (
	"mime/multipart"
	"time"
)

type Presence struct {
	ID        int       `json:"id" gorm:"primaryKey,autoIncrement"`
	UserID    int       `json:"user_id"`
	Location  string    `json:"location" binding:"required"`
	Type      string    `json:"type" binding:"required,oneof=masuk keluar izin"`
	Image     *string   `json:"image"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;<-:create"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	User *User `json:"user" gorm:"foreignKey:UserID"`
}

type PresenceRequest struct {
	UserID   int                   `form:"user_id" binding:"required"`
	Location string                `form:"location" binding:"required"`
	Type     string                `form:"type" binding:"required,oneof=masuk keluar izin"`
	Image    *multipart.FileHeader `form:"image"`
}

type PresenceUserRequest struct {
	Location string                `form:"location" binding:"required"`
	Type     string                `form:"type" binding:"required,oneof=masuk keluar izin"`
	Image    *multipart.FileHeader `form:"image"`
}
