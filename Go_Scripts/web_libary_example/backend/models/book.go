// models/book.go
package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Book struct {
	ID          ulid.ULID  `gorm:"primaryKey;type:char(26)"`
	Title       string     `gorm:"not null"`
	AuthorID    ulid.ULID  `gorm:"type:char(26);not null"`
	Author      Author     `gorm:"foreignKey:AuthorID"`
	Categories  []Category `gorm:"many2many:book_categories;"`
	ISBN        string     `gorm:"uniqueIndex;not null"`
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GenerateULID generates a new ULID for new records.
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = ulid.Make()
	return
}
