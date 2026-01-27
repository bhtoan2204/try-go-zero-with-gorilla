package repository

import (
	"context"
	appCtx "go-socket/core/context"
	"go-socket/core/domain/room/entity"
	"go-socket/core/domain/room/repos"
	"go-socket/core/infra/cache"
	"go-socket/core/infra/persistent/models"

	"gorm.io/gorm"
)

type roomRepoImpl struct {
	db        *gorm.DB
	roomCache *cache.RoomCache
}

func NewRoomRepoImpl(ctx context.Context, appCtx *appCtx.AppContext) repos.RoomRepository {
	return &roomRepoImpl{
		db:        appCtx.GetDB(),
		roomCache: cache.NewRoomCache(appCtx.GetCache()),
	}
}

func (r *roomRepoImpl) CreateRoom(ctx context.Context, room *entity.Room) error {
	return r.db.Create(room).Error
}

func (r *roomRepoImpl) GetRoomByID(ctx context.Context, id string) (*entity.Room, error) {
	if m, ok, err := r.roomCache.Get(ctx, id); err == nil && ok {
		return r.toEntity(m), nil
	}
	var m models.RoomModel
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error
	if err != nil {
		return nil, err
	}
	_ = r.roomCache.Set(ctx, &m)
	return r.toEntity(&m), nil
}

func (r *roomRepoImpl) UpdateRoom(ctx context.Context, room *entity.Room) error {
	return r.db.Save(room).Error
}

func (r *roomRepoImpl) DeleteRoom(ctx context.Context, id string) error {
	return r.db.Delete(&entity.Room{}, "id = ?", id).Error
}

func (r *roomRepoImpl) toEntity(m *models.RoomModel) *entity.Room {
	return &entity.Room{
		ID:        m.ID,
		Name:      m.Name,
		OwnerID:   m.OwnerID,
		OwnerType: m.OwnerType,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r *roomRepoImpl) toModel(e *entity.Room) *models.RoomModel {
	return &models.RoomModel{
		ID:        e.ID,
		Name:      e.Name,
		OwnerID:   e.OwnerID,
		OwnerType: e.OwnerType,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
