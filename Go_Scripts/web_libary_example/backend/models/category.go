// models/category.go
package models

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Category struct {
	ID        ulid.ULID `gorm:"primaryKey;type:char(26)"`
	Name      string    `gorm:"not null;unique"`
	Books     []Book    `gorm:"many2many:book_categories;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = ulid.Make()
	return
}
