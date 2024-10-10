package models

import "time"

type File struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"unique;size:100;not null"`
	Body      []byte    `json:"body" gorm:"type:longblob;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type FileResponse struct {
	File string `json:"file"`
}

type FileFilenameParam struct {
	Filename string `uri:"filename" binding:"required"`
}
