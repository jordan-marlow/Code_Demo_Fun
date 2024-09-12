package models

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Order struct {
	ID         string    `gorm:"type:char(26);primaryKey"`
	Customer   string    `gorm:"size:100;not null"`
	Products   []Product `gorm:"many2many:order_products;"` // Association with the Product model
	TotalCents int64     `gorm:"not null"`                  // To store total order value in cents
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (product *Order) BeforeCreate(tx *gorm.DB) (err error) {
	entropy := ulid.Monotonic(rand.Reader, 0)                               // Generate entropy
	product.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String() // Generate ULID
	return
}
