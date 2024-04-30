package services

import (
	clientproto "client-service/genproto/client_service"
	"client-service/internal/infrastructure/grpc_service_clients"
	"client-service/internal/pkg/otlp"
	"client-service/internal/usecase"
	"context"

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

	return nil, nil
}

func (s clientRPC) UpdateClient(ctx context.Context, in *clientproto.Client) (*clientproto.Client, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Id),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) DeleteClient(ctx context.Context, in *clientproto.ClientWithGUID) (*clientproto.DeleteClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Guid),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) GetClient(ctx context.Context, in *clientproto.ClientWithGUID) (*clientproto.Client, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetClient")
	span.SetAttributes(
		attribute.Key("guid").String(in.Guid),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) GetAllClients(ctx context.Context, in *clientproto.ListRequest) (*clientproto.ListClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) GetAllDeletedClients(ctx context.Context, in *clientproto.ListRequest) (*clientproto.ListClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetAllDeletedClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) GetAllHiddenClients(ctx context.Context, in *clientproto.ListRequest) (*clientproto.ListClientResponse, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetAllHiddenClients")
	span.SetAttributes(
		attribute.Key("guid").String(in.String()),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) UniqueEmail(ctx context.Context, in *clientproto.IsUnique) (*clientproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UniqueEmail")
	span.SetAttributes(
		attribute.Key("guid").String(in.Email),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) UpdateRefresh(ctx context.Context, in *clientproto.RefreshRequest) (*clientproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateRefresh")
	span.SetAttributes(
		attribute.Key("guid").String(in.ClientId),
	)
	defer span.End()

	return nil, nil
}

func (s clientRPC) UpdatePassword(ctx context.Context, in *clientproto.UpdatePasswordRequest) (*clientproto.ResponseStatus, error) {
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdatePassword")
	span.SetAttributes(
		attribute.Key("guid").String(in.ClientId),
	)
	defer span.End()

	return nil, nil
}
