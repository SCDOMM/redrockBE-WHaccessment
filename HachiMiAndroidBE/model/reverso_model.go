package model

import (
	"time"

	"gorm.io/gorm"
)

const DefaultImage string = "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAACAAQDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIRAxEAPwCdABmX/9k="

type UserModel struct {
	ID           uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	UserName     string `gorm:"unique;not null"`
	Account      string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Role         uint   `gorm:"default:0;not null"`
	ProfileImage string `gorm:"type:longtext;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	gorm.DeletedAt
}

type LogonDTO struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type LogonResponseDTO struct {
	UserName     string `json:"user_name"`
	ProfileImage string `json:"profile_image"`
	OperationDTO
}

type RegisterDTO struct {
	UserName string `json:"user_name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

type OperationDTO struct {
	ExpirationTime time.Time `json:"expiration_time"`
	UserRole       uint      `json:"user_role"`
	AccessToken    string    `json:"access_token"`
	RefreshToken   string    `json:"refresh_token"`
}
type ChangeProfileDTO struct {
	Name         string `json:"user_name"`
	Account      string `json:"account"`
	ProfileImage string `json:"profile_image"`
}
