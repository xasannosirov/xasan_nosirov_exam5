package repository

import (
	"client-service/internal/entity"
	"context"
)

type Clients interface {
	CreateClient(ctx context.Context, kyc *entity.Client) (*entity.Client, error)
	UpdateClient(ctx context.Context, kyc *entity.Client) (*entity.Client, error)
	DeleteClient(ctx context.Context, guid string) error
	GetClient(ctx context.Context, params map[string]string) (*entity.Client, error)
	GetAllClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Client, error)
	GetAllDeletedClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Client, error)
	GetAllHiddenClients(ctx context.Context, limit, offset uint64, filter map[string]bool) ([]*entity.Client, error)
	UniqueEmail(ctx context.Context, request *entity.IsUnique) (*entity.Response, error)
	UpdateRefresh(ctx context.Context, request *entity.UpdateRefresh) (*entity.Response, error)
	UpdatePassword(ctx context.Context, request *entity.UpdatePassword) (*entity.Response, error)
}
