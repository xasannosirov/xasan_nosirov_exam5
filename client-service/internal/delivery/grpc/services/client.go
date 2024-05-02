package services

import (
	clientproto "client-service/genproto/client_service"
	"client-service/internal/entity"
	"client-service/internal/infrastructure/grpc_service_clients"
	"client-service/internal/pkg/otlp"
	"client-service/internal/usecase"
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type clientRPC struct {
	logger        *zap.Logger
	clientUsecase usecase.Client
	clients       grpc_service_clients.ServiceClients
}

func NewRPC(logger *zap.Logger, clientUsecase usecase.Client, services *grpc_service_clients.ServiceClients) clientproto.ClientServiceServer {
	return &clientRPC{
		logger:        logger,
		clientUsecase: clientUsecase,
		clients:       *services,
	}
}

func (s clientRPC) CreateClient(ctx context.Context, in *clientproto.Client) (*clientproto.ClientWithGUID, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CreateClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Id),
	)
	defer span.End()

	createdClient, err := s.clientUsecase.CreateClient(ctx, &entity.Client{
		GUID:        in.Id,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Age:         uint64(in.Age),
		Gender:      in.Gender,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		Email:       in.Email,
		Password:    in.Password,
		Status:      in.Status,
		Refresh:     in.Refresh,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &clientproto.ClientWithGUID{
		Guid: createdClient.GUID,
	}, nil
}

func (s clientRPC) UpdateClient(ctx context.Context, in *clientproto.Client) (*clientproto.Client, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Id),
	)
	defer span.End()

	updatedClient, err := s.clientUsecase.UpdateClient(ctx, &entity.Client{
		GUID:        in.Id,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Age:         uint64(in.Age),
		Gender:      in.Gender,
		PhoneNumber: in.PhoneNumber,
		Address:     in.Address,
		Email:       in.Email,
		Password:    in.Password,
		Status:      in.Status,
		Refresh:     in.Refresh,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &clientproto.Client{
		Id:          updatedClient.GUID,
		FirstName:   updatedClient.FirstName,
		LastName:    updatedClient.LastName,
		Age:         uint32(updatedClient.Age),
		Gender:      updatedClient.Gender,
		PhoneNumber: updatedClient.PhoneNumber,
		Address:     updatedClient.Address,
		Email:       updatedClient.Email,
		Password:    updatedClient.Password,
		Status:      updatedClient.Status,
		Refresh:     updatedClient.Refresh,
		CreatedAt:   updatedClient.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updatedClient.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s clientRPC) DeleteClient(ctx context.Context, in *clientproto.ClientWithGUID) (*clientproto.DeleteClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Guid),
	)
	defer span.End()

	err := s.clientUsecase.DeleteClient(ctx, in.Guid)
	if err != nil {
		s.logger.Error(err.Error())
		return &clientproto.DeleteClientResponse{Status: false}, err
	}

	return &clientproto.DeleteClientResponse{Status: true}, nil
}

func (s clientRPC) GetClient(ctx context.Context, in *clientproto.ClientWithGUID) (*clientproto.Client, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Guid),
	)
	defer span.End()

	client, err := s.clientUsecase.GetClient(ctx, map[string]string{
		"id": in.Guid,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &clientproto.Client{
		Id:          client.GUID,
		FirstName:   client.FirstName,
		LastName:    client.LastName,
		Age:         uint32(client.Age),
		Gender:      client.Gender,
		PhoneNumber: client.PhoneNumber,
		Address:     client.Address,
		Email:       client.Email,
		Password:    client.Password,
		Status:      client.Status,
		Refresh:     client.Refresh,
		CreatedAt:   client.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   client.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s clientRPC) GetAllClients(ctx context.Context, in *clientproto.ListRequest) (*clientproto.ListClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	listClients, err := s.clientUsecase.GetAllClients(ctx, uint64(in.Limit), uint64(offset), map[string]string{})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response clientproto.ListClientResponse
	for _, client := range listClients {
		response.Clients = append(response.Clients, &clientproto.Client{
			Id:          client.GUID,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Age:         uint32(client.Age),
			Gender:      client.Gender,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
			Email:       client.Email,
			Password:    client.Password,
			Status:      client.Status,
			Refresh:     client.Refresh,
			CreatedAt:   client.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   client.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s clientRPC) GetAllDeletedClients(ctx context.Context, in *clientproto.ListRequest) (*clientproto.ListClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetAllDeletedClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	listDeletedClients, err := s.clientUsecase.GetAllDeletedClients(ctx, uint64(in.Limit), uint64(offset), map[string]string{})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response clientproto.ListClientResponse
	for _, client := range listDeletedClients {
		response.Clients = append(response.Clients, &clientproto.Client{
			Id:          client.GUID,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Age:         uint32(client.Age),
			Gender:      client.Gender,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
			Email:       client.Email,
			Password:    client.Password,
			Status:      client.Status,
			Refresh:     client.Refresh,
			CreatedAt:   client.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   client.UpdatedAt.Format(time.RFC3339),
			DeletedAt:   client.DeletedAt.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s clientRPC) GetAllHiddenClients(ctx context.Context, in *clientproto.ListRequest) (*clientproto.ListClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetAllHiddenClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	offset := in.Limit * (in.Page - 1)

	listHiddenClients, err := s.clientUsecase.GetAllHiddenClients(ctx, uint64(in.Limit), uint64(offset), map[string]bool{
		"status": false,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response clientproto.ListClientResponse
	for _, client := range listHiddenClients {
		response.Clients = append(response.Clients, &clientproto.Client{
			Id:          client.GUID,
			FirstName:   client.FirstName,
			LastName:    client.LastName,
			Age:         uint32(client.Age),
			Gender:      client.Gender,
			PhoneNumber: client.PhoneNumber,
			Address:     client.Address,
			Email:       client.Email,
			Password:    client.Password,
			Status:      client.Status,
			Refresh:     client.Refresh,
			CreatedAt:   client.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   client.UpdatedAt.Format(time.RFC3339),
			DeletedAt:   client.DeletedAt.Format(time.RFC3339),
		})
	}

	return &response, nil
}

func (s clientRPC) UniqueEmail(ctx context.Context, in *clientproto.IsUnique) (*clientproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UniqueEmail")
	span.SetAttributes(
		attribute.Key("guid").String(in.Email),
	)
	defer span.End()

	_, err := s.clientUsecase.UniqueEmail(ctx, &entity.IsUnique{
		Email: in.Email,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return &clientproto.ResponseStatus{Status: false}, errors.New("Email is already in use")
	}

	return &clientproto.ResponseStatus{Status: true}, nil
}

func (s clientRPC) UpdateRefresh(ctx context.Context, in *clientproto.RefreshRequest) (*clientproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateRefresh")
	span.SetAttributes(
		attribute.Key("guid").String(in.ClientId),
	)
	defer span.End()

	status, err := s.clientUsecase.UpdateRefresh(ctx, &entity.UpdateRefresh{
		ClientID:     in.ClientId,
		RefreshToken: in.RefreshToken,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return &clientproto.ResponseStatus{Status: false}, err
	}
	if !status.Status {
		s.logger.Error(err.Error())
		return &clientproto.ResponseStatus{Status: false}, errors.New("doesn't update refresh token")
	}

	return &clientproto.ResponseStatus{Status: true}, nil
}

func (s clientRPC) UpdatePassword(ctx context.Context, in *clientproto.UpdatePasswordRequest) (*clientproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdatePassword")
	span.SetAttributes(
		attribute.Key("guid").String(in.ClientId),
	)
	defer span.End()

	status, err := s.clientUsecase.UpdatePassword(ctx, &entity.UpdatePassword{
		ClientID:    in.ClientId,
		NewPassword: in.NewPassword,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return &clientproto.ResponseStatus{Status: false}, err
	}
	if !status.Status {
		s.logger.Error(err.Error())
		return &clientproto.ResponseStatus{Status: false}, errors.New("doesn't update password")
	}

	return &clientproto.ResponseStatus{Status: true}, nil
}
