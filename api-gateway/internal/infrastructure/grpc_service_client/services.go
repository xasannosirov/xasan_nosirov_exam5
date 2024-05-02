package grpc_service_clients

import (
	"api-gateway/internal/pkg/config"
	"fmt"

	clientproto "api-gateway/genproto/client_service"
	jobproto "api-gateway/genproto/job_service"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type ServiceClient interface {
	JobService() jobproto.JobServiceClient
	ClientService() clientproto.ClientServiceClient
	Close()
}

type serviceClient struct {
	connections   []*grpc.ClientConn
	clientService clientproto.ClientServiceClient
	jobService    jobproto.JobServiceClient
}

func New(cfg *config.Config) (ServiceClient, error) {
	clientServiceConnection, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ClientService.Host, cfg.ClientService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	jobServiceConnection, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.JobService.Host, cfg.JobService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClient{
		clientService: clientproto.NewClientServiceClient(clientServiceConnection),
		jobService:    jobproto.NewJobServiceClient(jobServiceConnection),
		connections: []*grpc.ClientConn{
			clientServiceConnection,
			jobServiceConnection,
		},
	}, nil
}

func (s *serviceClient) ClientService() clientproto.ClientServiceClient {
	return s.clientService
}

func (s *serviceClient) JobService() jobproto.JobServiceClient {
	return s.jobService
}

func (s *serviceClient) Close() {
	for _, conn := range s.connections {
		if err := conn.Close(); err != nil {
			// should be replaced by logger soon
			fmt.Printf("error while closing grpc connection: %v", err)
		}
	}
}
