package models

import (
	"math/rand"
	"time"
)

type Url struct {
	ID          int       `json:"id" gorm:"primary_key"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ViewCount   int       `json:"view_count" gorm:"default:0"`
}

type CreateUrlInput struct {
	OriginalUrl string `json:"link" binding:"required"`
}

func GenerateShortUrl() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
