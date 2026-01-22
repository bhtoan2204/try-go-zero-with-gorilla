package models

import (
	"go-socket/types"
	"time"
)

type RoomMemberModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	RoomID    string         `gorm:"not null;index"`
	AccountID string         `gorm:"not null;index"`
	Role      types.RoomRole `gorm:"default:member"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
}

func (RoomMemberModel) TableName() string {
	return "room_members"
}
