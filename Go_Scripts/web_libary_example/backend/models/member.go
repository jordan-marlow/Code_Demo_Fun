// models/member.go
package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Member struct {
	ID        ulid.ULID `gorm:"primaryKey;type:char(26)"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Member) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = ulid.Make()
	return
}
