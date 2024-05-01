package usecase

import (
	"context"
	"job-service/internal/entity"
	"job-service/internal/infrastructure/repository"
	"job-service/internal/pkg/otlp"
	"time"
)

type Job interface {
	CreateJob(ctx context.Context, article *entity.Job) (*entity.Job, error)
	UpdateJob(ctx context.Context, article *entity.Job) (*entity.Job, error)
	DeleteJob(ctx context.Context, guid string) error
	GetJob(ctx context.Context, params map[string]string) (*entity.Job, error)
	GetAllJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Job, error)
	GetAllDeletedJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Job, error)
	GetClientJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.ClientJob, error)
	GetJobClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.ClientJob, error)
	AddClientJob(ctx context.Context, job *entity.ClientJob) (*entity.Response, error)
	DeleteClientJob(ctx context.Context, clientJob *entity.ClientJob) error
}

type jobService struct {
	BaseUseCase
	repo       repository.Jobs
	ctxTimeout time.Duration
}

func NewJobService(ctxTimeout time.Duration, repo repository.Jobs) Job {
	return jobService{
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

func (u jobService) CreateJob(ctx context.Context, job *entity.Job) (*entity.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "CreateClient")
	defer span.End()

	u.beforeRequest(&job.GUID, &job.CreatedAt, &job.UpdatedAt)

	return u.repo.CreateJob(ctx, job)
}

func (u jobService) UpdateJob(ctx context.Context, job *entity.Job) (*entity.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "UpdateClient")
	defer span.End()

	u.beforeRequest(nil, nil, &job.UpdatedAt)

	return u.repo.UpdateJob(ctx, job)
}

func (u jobService) DeleteJob(ctx context.Context, guid string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "DeleteClient")
	defer span.End()

	return u.repo.DeleteJob(ctx, guid)
}

func (u jobService) GetJob(ctx context.Context, params map[string]string) (*entity.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "GetClient")
	defer span.End()

	return u.repo.GetJob(ctx, params)
}

func (u jobService) GetAllJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListClients")
	defer span.End()

	return u.repo.GetAllJobs(ctx, limit, offset, filter)
}

func (u jobService) GetAllDeletedJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.Job, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListDeletedClients")
	defer span.End()

	return u.repo.GetAllDeletedJobs(ctx, limit, offset, filter)
}

func (u jobService) GetClientJobs(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.ClientJob, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListHiddenUsers")
	defer span.End()

	return u.repo.GetClientJobs(ctx, limit, offset, filter)
}

func (u jobService) GetJobClients(ctx context.Context, limit, offset uint64, filter map[string]string) ([]*entity.ClientJob, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "ListHiddenUsers")
	defer span.End()

	return u.repo.GetClientJobs(ctx, limit, offset, filter)
}

func (u jobService) AddClientJob(ctx context.Context, clientJob *entity.ClientJob) (*entity.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "AddClientJob")
	span.End()

	return u.repo.AddClientJob(ctx, clientJob)
}

func (u jobService) DeleteClientJob(ctx context.Context, clientJob *entity.ClientJob) error {
	ctx, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	ctx, span := otlp.Start(ctx, "user_grpc-usercase", "DeleteClientJob")
	defer span.End()

	return u.repo.DeleteClientJob(ctx, clientJob)
}
