package lib

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID uuid.UUID `gorm:"type:varchar(36);primary_key;default:(uuid());not null"`
	gorm.Model
}

func (order *Model) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	order.ID = id
	return err
}