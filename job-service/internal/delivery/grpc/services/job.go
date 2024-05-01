package services

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	jobproto "job-service/genproto/job_service"
	"job-service/internal/entity"
	"job-service/internal/infrastructure/grpc_service_clients"
	"job-service/internal/pkg/otlp"
	"job-service/internal/usecase"
	"time"
)

type jobRPC struct {
	logger     *zap.Logger
	jobUsecase usecase.Job
	clients    grpc_service_clients.ServiceClients
}

func NewRPC(logger *zap.Logger, jobUsecase usecase.Job, services *grpc_service_clients.ServiceClients) jobproto.JobServiceServer {
	return &jobRPC{
		logger:     logger,
		jobUsecase: jobUsecase,
		clients:    *services,
	}
}

func (s jobRPC) CreateJob(ctx context.Context, in *jobproto.Job) (*jobproto.JobWithGUID, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CreateClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Id),
	)
	defer span.End()

	createdJob, err := s.jobUsecase.CreateJob(ctx, &entity.Job{
		GUID:           in.Id,
		Name:           in.Name,
		Salary:         float64(in.Salary),
		Level:          in.Level,
		LocationType:   in.LocationType,
		EmploymentType: in.EmploymentType,
		Address:        in.Address,
		Company:        in.Company,
	})
	if err != nil {
		return nil, err
	}

	return &jobproto.JobWithGUID{
		JobId: createdJob.GUID,
	}, nil
}

func (s jobRPC) UpdateJob(ctx context.Context, in *jobproto.Job) (*jobproto.Job, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Id),
	)
	defer span.End()

	createdAt, err := time.Parse(time.RFC3339, in.CreatedAt)
	if err != nil {
		return nil, err
	}
	updatedAt, err := time.Parse(time.RFC3339, in.UpdatedAt)
	if err != nil {
		return nil, err
	}

	updatedJob, err := s.jobUsecase.UpdateJob(ctx, &entity.Job{
		GUID:           in.Id,
		Name:           in.Name,
		Salary:         float64(in.Salary),
		Level:          in.Level,
		LocationType:   in.LocationType,
		EmploymentType: in.EmploymentType,
		Address:        in.Address,
		UpdatedAt:      updatedAt,
		CreatedAt:      createdAt,
	})
	if err != nil {
		return nil, err
	}

	return &jobproto.Job{
		Id:             updatedJob.GUID,
		Name:           updatedJob.Name,
		Salary:         float32(updatedJob.Salary),
		Level:          updatedJob.Level,
		LocationType:   updatedJob.LocationType,
		EmploymentType: updatedJob.EmploymentType,
		Address:        updatedJob.Address,
		Company:        updatedJob.Company,
		CreatedAt:      updatedJob.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      updatedJob.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s jobRPC) DeleteJob(ctx context.Context, in *jobproto.JobWithGUID) (*jobproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.JobId),
	)
	defer span.End()

	err := s.jobUsecase.DeleteJob(ctx, in.JobId)
	if err != nil {
		return &jobproto.ResponseStatus{Status: false}, err
	}

	return &jobproto.ResponseStatus{Status: true}, nil
}

func (s jobRPC) GetJob(ctx context.Context, in *jobproto.JobWithGUID) (*jobproto.Job, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.JobId),
	)
	defer span.End()

	job, err := s.jobUsecase.GetJob(ctx, map[string]string{
		"id": in.JobId,
	})
	if err != nil {
		return nil, err
	}

	return &jobproto.Job{
		Id:             job.GUID,
		Name:           job.Name,
		Salary:         float32(job.Salary),
		Level:          job.Level,
		LocationType:   job.LocationType,
		EmploymentType: job.EmploymentType,
		Address:        job.Address,
		Company:        job.Company,
		CreatedAt:      job.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      job.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s jobRPC) GetAllJobs(ctx context.Context, in *jobproto.ListRequest) (*jobproto.ListJobResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	listJobs, err := s.jobUsecase.GetAllJobs(ctx, in.Limit, offset, map[string]string{})
	if err != nil {
		return nil, err
	}

	var response jobproto.ListJobResponse
	for _, job := range listJobs {
		response.Jobs = append(response.Jobs, &jobproto.Job{
			Id:             job.GUID,
			Name:           job.Name,
			Salary:         float32(job.Salary),
			Level:          job.Level,
			LocationType:   job.LocationType,
			EmploymentType: job.EmploymentType,
			Address:        job.Address,
			Company:        job.Company,
			CreatedAt:      job.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      job.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s jobRPC) GetAllDeletedJobs(ctx context.Context, in *jobproto.ListRequest) (*jobproto.ListJobResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetAllDeletedClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	listDeletedJobs, err := s.jobUsecase.GetAllDeletedJobs(ctx, in.Limit, offset, map[string]string{})
	if err != nil {
		return nil, err
	}

	var response jobproto.ListJobResponse
	for _, job := range listDeletedJobs {
		response.Jobs = append(response.Jobs, &jobproto.Job{
			Id:             job.GUID,
			Name:           job.Name,
			Salary:         float32(job.Salary),
			Level:          job.Level,
			EmploymentType: job.EmploymentType,
			Address:        job.Address,
			Company:        job.Company,
			CreatedAt:      job.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      job.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s jobRPC) GetClientJobs(ctx context.Context, in *jobproto.ClientJobRequest) (*jobproto.ListClientJobs, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetClientJobs")
	span.SetAttributes(
		attribute.Key("guid").String(in.JobId),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	clientJobs, err := s.jobUsecase.GetClientJobs(ctx, in.Limit, offset, map[string]string{
		"client_id": in.ClientId,
	})
	if err != nil {
		return nil, err
	}

	var response jobproto.ListClientJobs
	for _, clientJob := range clientJobs {
		response.ClientJobs = append(response.ClientJobs, &jobproto.ClientJobs{
			ClientId:  clientJob.ClientID,
			JobId:     clientJob.JobID,
			StartDate: clientJob.StartDate.Format(time.RFC3339),
			EndDate:   clientJob.EndDate.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s jobRPC) GetJobClients(ctx context.Context, in *jobproto.ClientJobRequest) (*jobproto.ListClientJobs, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetJobClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.JobId),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	clientJobs, err := s.jobUsecase.GetClientJobs(ctx, in.Limit, offset, map[string]string{
		"job_id": in.JobId,
	})
	if err != nil {
		return nil, err
	}

	var response jobproto.ListClientJobs
	for _, clientJob := range clientJobs {
		response.ClientJobs = append(response.ClientJobs, &jobproto.ClientJobs{
			ClientId:  clientJob.ClientID,
			JobId:     clientJob.JobID,
			StartDate: clientJob.StartDate.Format(time.RFC3339),
			EndDate:   clientJob.EndDate.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s jobRPC) AddClientJob(ctx context.Context, in *jobproto.ClientJobs) (*jobproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "AddClientJob")
	span.SetAttributes(
		attribute.Key("guid").String(in.JobId),
	)
	defer span.End()

	startDate, err := time.Parse(time.RFC3339, in.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse(time.RFC3339, in.EndDate)
	if err != nil {
		return nil, err
	}

	_, err = s.jobUsecase.AddClientJob(ctx, &entity.ClientJob{
		ClientID:  in.ClientId,
		JobID:     in.JobId,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return &jobproto.ResponseStatus{Status: false}, err
	}

	return &jobproto.ResponseStatus{Status: true}, nil
}

func (s jobRPC) DeleteClientJob(ctx context.Context, in *jobproto.ClientJobs) (*jobproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteClientJob")
	span.SetAttributes(
		attribute.Key("guid").String(in.JobId),
	)
	defer span.End()

	err := s.jobUsecase.DeleteClientJob(ctx, &entity.ClientJob{
		ClientID: in.ClientId,
		JobID:    in.JobId,
	})
	if err != nil {
		return &jobproto.ResponseStatus{Status: false}, err
	}

	return &jobproto.ResponseStatus{Status: true}, nil
}
