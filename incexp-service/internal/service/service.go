package service

import (
	"context"
	"fmt"
	pb "incexp-service/genproto/incexpb"
	producer "incexp-service/internal/kafka/producer"
	"incexp-service/internal/service/helper"
	"incexp-service/internal/storage"
	"incexp-service/logger"

	"github.com/google/uuid"
)

type IncExpService struct {
	pb.UnimplementedIncExpServiceServer
	Queries  *storage.Queries
	Producer producer.ProducerInit
}

func NewIncExpService(repo *storage.Queries, producer producer.ProducerInit) *IncExpService {
	return &IncExpService{
		Queries:  repo,
		Producer: producer,
	}
}

func (s *IncExpService) RegisterIncome(ctx context.Context, req *pb.RegisterIncomeRequest) (*pb.RegisterIncomeResponse, error) {

	userId, err := helper.ParseUUID(req.UserId)
	if err != nil {
		return nil, err
	}
	logger.Info("RegisterIncome: User ID parsed - ", userId)

	categoryId, err := helper.ParseUUID(req.CategoryId)
	if err != nil {
		return nil, err
	}
	logger.Info("RegisterIncome: Category ID parsed - ", categoryId)

	income := storage.RegisterIncomeParams{
		UserID:     userId,
		Type:       "income",
		Amount:     req.Amount,
		Currency:   req.Currency,
		CategoryID: uuid.NullUUID{UUID: categoryId, Valid: true},
	}

	id, err := s.Queries.RegisterIncome(ctx, income)
	if err != nil {
		logger.Error("RegisterIncome: Error registering income - ", err)
		return nil, err
	}

	resp := pb.RegisterIncomeResponse{
		Message:       "Income logged successfully",
		TransactionId: id.String(),
	}

	logger.Info("RegisterIncome: Income registered successfully - Transaction ID: ", resp.TransactionId)
	return &resp, nil
}

func (s *IncExpService) RegisterExpense(ctx context.Context, req *pb.RegisterExpenseRequest) (*pb.RegisterExpenseResponse, error) {

	userId, err := helper.ParseUUID(req.UserId)
	if err != nil {
		return nil, err
	}

	categoryId, err := helper.ParseUUID(req.CategoryId)
	if err != nil {
		return nil, err
	}

	logger.Info("RegisterExpense: Category ID parsed - ", categoryId)

	expense := storage.RegisterExpenseParams{
		UserID:     userId,
		Type:       "expense",
		Amount:     req.Amount,
		Currency:   req.Currency,
		CategoryID: uuid.NullUUID{UUID: categoryId, Valid: true},
	}

	id, err := s.Queries.RegisterExpense(ctx, expense)
	if err != nil {
		logger.Error("RegisterExpense: Error registering expense - ", err)
		return nil, err
	}
	params := storage.CheckBudgetParams{
		UserID:     userId,
		CategoryID: uuid.NullUUID{UUID: categoryId, Valid: true},
	}

	budgets, err := s.Queries.CheckBudget(ctx, params)
	if err != nil {
		return nil, err
	}

	if budgets.Status == "Over Budget" {
		message := fmt.Sprintf("User %s has exceeded the budget for category %v. Spent: %.2f, Budget: %d",
			req.UserEmail, budgets.CategoryID, budgets.TotalSpent, budgets.BudgetLimit)
		err = s.Producer.ProduceMessage(req.UserEmail, []byte(message))
		if err != nil {
			logger.Error("RegisterExpense: Error sending message to Kafka - ", err)
			return nil, err
		}
		logger.Info("RegisterExpense: Over budget message sent to Kafka for user - ", userId)
	}

	resp := pb.RegisterExpenseResponse{
		Message:       "Expense logged successfully",
		TransactionId: id.String(),
	}

	logger.Info("RegisterExpense: Expense registered successfully - Transaction ID: ", resp.TransactionId)
	return &resp, nil
}

func (s *IncExpService) GetListIncomeVSExpense(ctx context.Context, req *pb.GetListIncomeVSExpenseRequest) (*pb.GetListIncomeVSExpenseResponse, error) {

	logger.Info("GetListIncomeVSExpense: Fetching income vs expense list")

	userId, err := helper.ParseUUID(req.UserId)
	if err != nil {
		return nil, err
	}
	resp := pb.GetListIncomeVSExpenseResponse{}
	list, err := s.Queries.GetListIncomeVSExpense(ctx, userId)
	if err != nil {
		logger.Error("GetListIncomeVSExpense: Error retrieving income vs expense list - ", err)
		return nil, err
	}

	logger.Info("GetListIncomeVSExpense: Successfully retrieved income vs expense list - Count: ", len(list))

	for _, v := range list {
		item := &pb.IncomeVSExpense{
			TransactionId: v.ID.String(),
			Type:          string(v.Type),
			Amount:        v.Amount,
			Currency:      v.Currency,
			Category:      v.Category,
			Date:          v.Date.Time.Format("2006-01-02 15:04:05"),
		}
		resp.List = append(resp.List, item)
	}

	return &resp, nil
}
