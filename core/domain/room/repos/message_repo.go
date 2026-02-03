package repos

import (
	"context"
	"go-socket/core/domain/room/entity"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *entity.MessageEntity) error
}
