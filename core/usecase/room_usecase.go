package usecase

import (
	"context"
	"go-socket/core/domain/room/entity"
)

type RoomUsecase interface {
	CreateRoom(ctx context.Context, room *entity.Room) error
	GetRoomByID(ctx context.Context, id string) (*entity.Room, error)
	UpdateRoom(ctx context.Context, room *entity.Room) error
	DeleteRoom(ctx context.Context, id string) error
}
