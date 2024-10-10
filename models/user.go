package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey,autoIncrement"`
	Name      string    `json:"name"`
	Username  string    `json:"username" gorm:"unique;not null" binding:"required"`
	Password  string    `json:"password" gorm:"size:100;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;<-:create"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Presences *[]Presence `json:"presences" gorm:"foreignKey:UserID"`
}
