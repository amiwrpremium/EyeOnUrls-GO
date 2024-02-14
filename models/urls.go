package models

import (
	"math/rand"
	"time"
)

// Url represents the model for a shortened URL
type Url struct {
	ID          int       `json:"id" gorm:"primary_key"`
	OriginalUrl string    `json:"original_url"`
	ShortUrl    string    `json:"short_url" gorm:"unique"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	ViewCount   int       `json:"view_count" gorm:"default:0"`
}

// CreateUrlInput represents the input for the CreateUrl function
type CreateUrlInput struct {
	OriginalUrl string `json:"link" binding:"required"`
}

// GenerateShortUrl generates a random short URL
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
