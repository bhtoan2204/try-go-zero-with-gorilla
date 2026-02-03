package aggregate

import (
	"go-socket/core/domain/room/entity"
	"go-socket/types"
)

type RoomAggregate struct {
	RoomID   string
	RoomType types.RoomType
	Messages []*entity.MessageEntity
}
