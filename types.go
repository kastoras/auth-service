package main

import (
	"time"

	"gorm.io/gorm"
)

type Folder struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePostPayload struct {
	Title string `json:"title"`
}

type Storage struct {
	db *gorm.DB
}
