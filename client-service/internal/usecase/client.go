package usecase

import (
	"client-service/internal/entity"
	"client-service/internal/infrastructure/repository"
	"client-service/internal/pkg/otlp"
	"context"
	"time"
)

type Client interface {
	CreateClient(ctx context.Context, article *entity.Client) (*entity.Client, error)
	UpdateClient(ctx context.Context, article *entity.Client) (*entity.Client, error)
	DeleteClient(ctx context.Context, guid string) error
	GetClient(ctx context.Context, params map[string]string) (*entity.Client, error)
	GetAllClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Client, error)
	GetAllDeletedClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Client, error)
	GetAllHiddenClients(ctx context.Context, limit, offset uint64, filter map[string]bool) ([]*entity.Client, error)
	UniqueEmail(ctx context.Context, request *entity.IsUnique) (*entity.Response, error)
	UpdateRefresh(ctx context.Context, request *entity.UpdateRefresh) (*entity.Response, error)
	UpdatePassword(ctx context.Context, request *entity.UpdatePassword) (*entity.Response, error)
}

type clientService struct {
	BaseUseCase
	repo       repository.Clients
	ctxTimeout time.Duration
}

func NewUserService(ctxTimeout time.Duration, repo repository.Clients) Client {
	return clientService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u clientService) CreateClient(ctx context.Context, user *entity.Client) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "CreateClient")
	defer span.End()

	u.beforeRequest(&user.GUID, &user.CreatedAt, &user.UpdatedAt)

	return u.repo.CreateClient(ctx, user)
}

func (u clientService) UpdateClient(ctx context.Context, user *entity.Client) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "UpdateClient")
	defer span.End()

	u.beforeRequest(nil, nil, &user.UpdatedAt)

	return u.repo.UpdateClient(ctx, user)
}

func (u clientService) DeleteClient(ctx context.Context, guid string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "DeleteClient")
	defer span.End()

	return u.repo.DeleteClient(ctx, guid)
}

func (u clientService) GetClient(ctx context.Context, params map[string]string) (*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "GetClient")
	defer span.End()

	return u.repo.GetClient(ctx, params)
}

func (u clientService) GetAllClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListClients")
	defer span.End()

	return u.repo.GetAllClients(ctx, limit, offset, filter)
}

func (u clientService) GetAllDeletedClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListDeletedClients")
	defer span.End()

	return u.repo.GetAllDeletedClients(ctx, limit, offset, filter)
}

func (u clientService) GetAllHiddenClients(ctx context.Context, limit, offset uint64, filter map[string]bool) ([]*entity.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListHiddenUsers")
	defer span.End()

	filter["status"] = false
	return u.repo.GetAllHiddenClients(ctx, limit, offset, filter)
}

func (u clientService) UniqueEmail(ctx context.Context, request *entity.IsUnique) (*entity.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "UniqueEmail")
	defer span.End()

	return u.repo.UniqueEmail(ctx, request)
}

func (u clientService) UpdateRefresh(ctx context.Context, request *entity.UpdateRefresh) (*entity.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "UpdateRefresh")
	defer span.End()

	return u.repo.UpdateRefresh(ctx, request)
}

func (u clientService) UpdatePassword(ctx context.Context, request *entity.UpdatePassword) (*entity.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "UpdatePassword")
	defer span.End()

	return u.repo.UpdatePassword(ctx, request)
}
