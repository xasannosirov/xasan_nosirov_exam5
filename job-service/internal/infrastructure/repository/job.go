package repository

import (
	"context"
	"job-service/internal/entity"
)

type Jobs interface {
	CreateJob(ctx context.Context, kyc *entity.Job) (*entity.Job, error)
	UpdateJob(ctx context.Context, kyc *entity.Job) (*entity.Job, error)
	DeleteJob(ctx context.Context, guid string) error
	GetJob(ctx context.Context, params map[string]string) (*entity.Job, error)
	GetAllJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Job, error)
	GetAllDeletedJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Job, error)
	GetClientJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.ClientJob, error)
	GetJobClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.ClientJob, error)
	AddClientJob(ctx context.Context, kyc *entity.ClientJob) (*entity.Response, error)
	DeleteClientJob(ctx context.Context, clientJob *entity.ClientJob) error
}
