package models

type User struct {
	ID           int    `json:"id" gorm:"primary_key;autoIncrement"`
	Name         string `json:"name"`
	Email        string `json:"email" binding:"required"`
	PasswordHash string `json:"-"`
	AccessToken  string `json:"-"`
	RefreshToken string `json:"-"`
	CreatedAt    int    `json:"created_at" gorm:"autoCreateTime;<-:create"`
	UpdatedAt    int    `json:"updated_at" gorm:"autoUpdateTime"`
	CreatedBy    *int   `json:"created_by" gorm:"<-:create"`
	UpdatedBy    *int   `json:"updated_by"`
}
