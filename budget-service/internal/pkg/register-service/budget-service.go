package registerservice

import (
	"fmt"
	"net"
	config "budget-service/internal/pkg/load"
	"budget-service/internal/service"
	"budget-service/logger"

	pb "budget-service/genproto/budgetpb"

	"google.golang.org/grpc"
)

type Service struct {
	Service *service.BudgetService
}

func NewGRPCBudgetService(service *service.BudgetService) *Service {
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

	logger.Info("Budget service listening on :", cfg.Service.Port)

	pb.RegisterBudgetServiceServer(grpcServer, s.Service)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}

func GracefufShutdown() {
	grpcServer.GracefulStop()
}
