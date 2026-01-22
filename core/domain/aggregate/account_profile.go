package aggregate

import "go-socket/core/domain/entity"

type AccountProfile struct {
	entity.Account
	CreatedRooms []entity.Room
	JoinedRooms  []entity.Room
}
