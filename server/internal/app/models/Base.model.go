package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)
type Base struct {
	ID uuid.UUID `gorm:"uniqueIndex;type:char(36);primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
  b.ID  = uuid.New()
  return
}