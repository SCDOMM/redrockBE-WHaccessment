package model

import (
	"time"

	"gorm.io/gorm"
)

type HomeModel struct {
	ID        uint   `gorm:"primary_key;auto_increment;not_null;unique"`
	Title     string `gorm:"not null"`
	Desc      string `gorm:"not null"`
	Image     string `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
}
type HomeDTO struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
}
type SearchDTO struct {
	Content string `json:"content"`
}
type DynamicModel struct {
	ID            uint   `gorm:"primary_key;auto_increment;not_null;unique"`
	AuthorAccount string `gorm:"not null"`
	Title         string `gorm:"not null"`
	Desc          string `gorm:"not null"`
	CoverImage    string `gorm:"not null;type:longtext"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	gorm.DeletedAt
}
type DynamicDTO struct {
	Title        string    `json:"title"`
	Desc         string    `json:"desc"`
	AuthorName   string    `json:"author_name"`
	ProfileImage string    `json:"profile_image"`
	CoverImage   string    `json:"cover_image"`
	Account      string    `json:"account"`
	Time         time.Time `json:"time"`
}

type DynamicUploadDTO struct {
	AuthorAccount string `json:"author_account"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImage    string `json:"cover_image"`
}
