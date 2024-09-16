package registerservice

import (
	"fmt"
	"net"
	config "user-service/internal/pkg/load"
	"user-service/internal/service"
	"user-service/logger"

	pb "user-service/genproto/userpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.UserService
}

func NewGRPCUserService(service *service.UserService) *Service {
	return &Service{
		Service: service,
	}
}

var grpcServer = grpc.NewServer()

func (s *Service) RUN(cfg *config.Config) error {

	addr := fmt.Sprintf(":%s", cfg.Service.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Info("User-service listening on :", cfg.Service.Port)

	pb.RegisterUserServiceServer(grpcServer, s.Service)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}
