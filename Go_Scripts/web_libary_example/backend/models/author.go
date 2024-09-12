// models/author.go
package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Author struct {
	ID        ulid.ULID `gorm:"primaryKey;type:char(26)"`
	Name      string    `gorm:"not null"`
	Books     []Book
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = ulid.Make()
	return
}
