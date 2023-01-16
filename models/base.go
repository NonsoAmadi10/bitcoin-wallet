package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base Struct will contain columns that cut across all tables
type Model struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

// Set Primary Key ID as UUID while saving to database
func (b *Model) BeforeCreate(tx *gorm.DB) error {

	u := uuid.New()

	b.ID = u
	return nil
}
