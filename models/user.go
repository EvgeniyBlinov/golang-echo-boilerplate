package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID     uuid.UUID `json:"uuid" gorm:"type:uuid;default:uuid_generate_v7();primaryKey"`
	Email    string    `json:"email" gorm:"type:varchar(200);"`
	Name     string    `json:"name" gorm:"type:varchar(200);"`
	Password string    `json:"password" gorm:"type:varchar(200);"`
	//Posts    []Post    `json:",omitempty" gorm:"ForeignKey:UserUUID;AssociationForeignKey:UUID"`
	Posts []Post `json:",omitempty" gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
