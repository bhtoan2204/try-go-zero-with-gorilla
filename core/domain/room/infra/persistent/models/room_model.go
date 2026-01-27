package models

import "time"

type RoomModel struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	OwnerID   string    `gorm:"not null"`
	OwnerType string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (RoomModel) TableName() string {
	return "rooms"
}
