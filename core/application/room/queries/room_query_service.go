package queries

import (
	"context"
	"go-socket/core/domain/room/entity"
)

type RoomQueryService interface {
	GetRoom(ctx context.Context, id string) (*entity.Room, error)
}
