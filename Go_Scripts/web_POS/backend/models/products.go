package models

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Product struct {
	ID         string `gorm:"type:char(26);primaryKey"`
	Name       string `gorm:"size:100;not null"`
	PriceCents int64  `gorm:"not null"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.Reader, 0)                               // Generate entropy
	product.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String() // Generate ULID
	return
}
