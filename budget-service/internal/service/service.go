package service

import (
	pb "budget-service/genproto/budgetpb"
	"budget-service/internal/service/helper"
	"budget-service/internal/storage"
	"budget-service/logger"
	"context"

	"github.com/google/uuid"
)

type BudgetService struct {
	pb.UnimplementedBudgetServiceServer
	Queries *storage.Queries
}

func NewBudgetService(repo *storage.Queries) *BudgetService {
	return &BudgetService{
		Queries: repo,
	}
}

func (s *BudgetService) CreateBudget(ctx context.Context, req *pb.CreateBudgetRequest) (*pb.CreateBudgetResponse, error) {

	logger.Info("CreateBudget: Creating budget - CategoryId: ", req.CategoryId, ", Amount: ", req.Amount, ", Currency: ", req.Currency)

	userId, err := helper.ParseUUID(req.UserId)
	if err != nil {
		return nil, err
	}

	categoryId, err := helper.ParseUUID(req.CategoryId)
	if err != nil {
		return nil, err
	}

	budget := storage.CreateBudgetParams{
		UserID:     userId,
		CategoryID: uuid.NullUUID{UUID: categoryId, Valid: true},
		Amount:     req.Amount,
		Currency:   req.Currency,
	}

	id, err := s.Queries.CreateBudget(ctx, budget)
	if err != nil {
		logger.Error("CreateBudget: Error creating budget - ", err)
		return nil, err
	}

	resp := pb.CreateBudgetResponse{
		Message:  "Budget created succesfully",
		BudgetId: id.String(),
	}

	logger.Info("CreateBudget: Budget created successfully - BudgetId: ", resp.BudgetId)
	return &resp, nil
}

func (s *BudgetService) GetListBudget(ctx context.Context, req *pb.GetListBudgetRequest) (*pb.GetListBudgetResponse, error) {

	logger.Info("GetListBudget: Fetching list of budgets")
	resp := pb.GetListBudgetResponse{}

	userId, err := helper.ParseUUID(req.UserId)
	if err != nil {
		return nil, err
	}
	
	budgets, err := s.Queries.GetListBudget(ctx, userId)
	if err != nil {
		logger.Error("GetListBudget: Error retrieving budgets list - ", err)
		return nil, err
	}

	for _, budget := range budgets {
		item := &pb.Budget{
			BudgetId: budget.BudgetID.String(),
			Category: budget.Category,
			Amount:   budget.Amount,
			Spent:    budget.Spent,
			Currency: budget.Currency,
		}
		resp.List = append(resp.List, item)
	}

	logger.Info("GetListBudget: Successfully retrieved budgets list - Count: ", len(resp.List))
	return &resp, nil
}

func (s *BudgetService) UpdateBudget(ctx context.Context, req *pb.UpdateBudgetRequest) (*pb.UpdateBudgetResponse, error) {

	logger.Info("UpdateBudget: Updating budget - BudgetId: ", req.BudgetId, ", Amount: ", req.Amount)

	userId, err := helper.ParseUUID(req.UserId)
	if err != nil {
		return nil, err
	}

	id, err := helper.ParseUUID(req.BudgetId)
	if err != nil {
		return nil, err
	}

	budget := storage.UpdateBudgetParams{
		ID:     id,
		Amount: req.Amount,
		UserID: userId,
	}

	message, err := s.Queries.UpdateBudget(ctx, budget)
	if err != nil {
		logger.Error("UpdateBudget: Error updating budget - ", err)
		return nil, err
	}
	resp := pb.UpdateBudgetResponse{
		Message: message,
	}

	logger.Info("UpdateBudget: Budget updated successfully - BudgetId: ", req.BudgetId)
	return &resp, nil
}
