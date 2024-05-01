package grpc_service_clients

import (
	"fmt"
	clientproto "job-service/genproto/client_service"
	"job-service/internal/pkg/config"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type ServiceClients interface {
	ClientService() clientproto.ClientServiceClient
	Close()
}

type serviceClients struct {
	clientService clientproto.ClientServiceClient
	services      []*grpc.ClientConn
}

func New(config *config.Config) (ServiceClients, error) {

	clientServiceConnection, err := grpc.Dial(
		fmt.Sprintf("%s%s", config.ClientService.Host, config.ClientService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClients{
		clientService: clientproto.NewClientServiceClient(clientServiceConnection),
		services:      []*grpc.ClientConn{},
	}, nil
}

func (s *serviceClients) Close() {
	// closing job-client services
	for _, conn := range s.services {
		conn.Close()
	}
}

func (s *serviceClients) ClientService() clientproto.ClientServiceClient {
	return s.clientService
}
