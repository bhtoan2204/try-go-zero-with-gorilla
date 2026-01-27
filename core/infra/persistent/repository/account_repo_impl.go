package repository

import (
	"context"
	appCtx "go-socket/core/context"
	"go-socket/core/domain/account/entity"
	"go-socket/core/domain/account/repos"
	"go-socket/core/infra/cache"
	"go-socket/core/infra/persistent/models"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

type accountRepoImpl struct {
	db           *gorm.DB
	accountCache *cache.AccountCache
}

func NewAccountRepoImpl(ctx context.Context, appCtx *appCtx.AppContext) repos.AccountRepository {
	return &accountRepoImpl{
		db:           appCtx.GetDB(),
		accountCache: cache.NewAccountCache(appCtx.GetCache()),
	}
}

func (r *accountRepoImpl) GetAccountByID(ctx context.Context, id string) (*entity.Account, error) {
	if m, ok, err := r.accountCache.Get(ctx, id); err == nil && ok {
		return r.toEntity(m), nil
	}
	var m models.AccountModel

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&m).Error

	if err != nil {
		return nil, err
	}

	_ = r.accountCache.Set(ctx, &m)

	return r.toEntity(&m), nil
}

func (r *accountRepoImpl) GetAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var m models.AccountModel
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&m).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repos.ErrAccountNotFound
		}
		return nil, err
	}
	_ = r.accountCache.Set(ctx, &m)
	return r.toEntity(&m), nil
}

func (r *accountRepoImpl) CreateAccount(ctx context.Context, account *entity.Account) error {
	m := r.toModel(account)

	err := r.db.WithContext(ctx).
		Create(m).Error
	if err != nil {
		return err
	}

	return r.accountCache.Set(ctx, m)
}

func (r *accountRepoImpl) UpdateAccount(ctx context.Context, account *entity.Account) error {
	m := r.toModel(account)

	err := r.db.WithContext(ctx).
		Save(m).Error
	if err != nil {
		return err
	}

	return r.accountCache.Set(ctx, m)
}

func (r *accountRepoImpl) DeleteAccount(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).
		Delete(&models.AccountModel{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return r.accountCache.Delete(ctx, id)
}

func (r *accountRepoImpl) ListAccountsByRoomID(ctx context.Context, roomID string) ([]*entity.Account, error) {
	var accounts []*models.AccountModel
	err := r.db.WithContext(ctx).
		Model(&models.AccountModel{}).
		Select("accounts.*").
		Joins("JOIN room_members rm ON rm.account_id = accounts.id").
		Where("rm.room_id = ?", roomID).
		Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return lo.Map(accounts, func(account *models.AccountModel, _ int) *entity.Account {
		return r.toEntity(account)
	}), nil
}

func (r *accountRepoImpl) toEntity(m *models.AccountModel) *entity.Account {
	return &entity.Account{
		ID:        m.ID,
		Email:     m.Email,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (r *accountRepoImpl) toModel(e *entity.Account) *models.AccountModel {
	return &models.AccountModel{
		ID:        e.ID,
		Email:     e.Email,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
