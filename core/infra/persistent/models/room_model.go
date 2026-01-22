package models

import "time"

type RoomModel struct {
	ID        string        `gorm:"primaryKey"`
	Name      string        `gorm:"not null"`
	OwnerID   string        `gorm:"not null"`
	OwnerType string        `gorm:"not null"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
	Owner     *AccountModel `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (RoomModel) TableName() string {
	return "rooms"
}
