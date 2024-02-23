package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	UUID     uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v7();primaryKey"`
	Title    string    `json:"title" gorm:"type:text"`
	Content  string    `json:"content" gorm:"type:text"`
	UserUUID uuid.UUID `json:"user_uuid" gorm:"column:user_uuid;type:uuid"`
	User     User      `json:"-" gorm:"ForeignKey:UserUUID;AssociationForeignKey:UUID;references:UUID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
