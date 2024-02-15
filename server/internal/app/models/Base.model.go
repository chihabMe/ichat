package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Base struct {
	ID uuid.UUID `gorm:"uniqueIndex;type:char(36);primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
  b.ID  = uuid.New()
  return
}