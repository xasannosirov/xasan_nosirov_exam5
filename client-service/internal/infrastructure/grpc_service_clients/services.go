package grpc_service_clients

import (
	jobproto "client-service/genproto/job_service"
	"client-service/internal/pkg/config"
	"fmt"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type ServiceClients interface {
	JobService() jobproto.JobServiceClient
	Close()
}

type serviceClients struct {
	jobsService jobproto.JobServiceClient
	services    []*grpc.ClientConn
}

func New(config *config.Config) (ServiceClients, error) {

	jobServiceConnection, err := grpc.Dial(
		fmt.Sprintf("%s%s", config.JobService.Host, config.JobService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClients{
		jobsService: jobproto.NewJobServiceClient(jobServiceConnection),
		services:    []*grpc.ClientConn{},
	}, nil
}

func (s *serviceClients) Close() {
	// closing job-client services
	for _, conn := range s.services {
		conn.Close()
	}
}

func (s *serviceClients) JobService() jobproto.JobServiceClient {
	return s.jobsService
}
