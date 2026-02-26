package user

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UUID         string    `gorm:"type:char(36);uniqueIndex" json:"uuid"`
	Email        string    `gorm:"type:varchar(191);uniqueIndex" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255)" json:"-"`
	Nickname     string    `gorm:"type:varchar(64)" json:"nickname"`
	AvatarURL    string    `gorm:"type:varchar(255)" json:"avatarUrl"`
	Bio          string    `gorm:"type:varchar(255)" json:"bio"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
