package repos

import (
	"context"
	"go-socket/core/domain/room/entity"
)

type RoomMemberRepository interface {
	CreateRoomMember(ctx context.Context, roomMember *entity.RoomMemberEntity) error
}
