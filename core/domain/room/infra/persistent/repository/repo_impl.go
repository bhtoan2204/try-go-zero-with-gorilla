package repository

import (
	appCtx "go-socket/core/context"
	"go-socket/core/domain/room/repos"
)

type repoImpl struct {
	roomRepo       repos.RoomRepository
	messageRepo    repos.MessageRepository
	roomMemberRepo repos.RoomMemberRepository
}

func NewRepoImpl(appCtx *appCtx.AppContext) repos.Repos {
	roomRepo := NewRoomRepoImpl(appCtx.GetDB(), appCtx.GetCache())
	return &repoImpl{roomRepo: roomRepo}
}

func (r *repoImpl) RoomRepository() repos.RoomRepository {
	return r.roomRepo
}

func (r *repoImpl) MessageRepository() repos.MessageRepository {
	return r.messageRepo
}

func (r *repoImpl) RoomMemberRepository() repos.RoomMemberRepository {
	return r.roomMemberRepo
}
