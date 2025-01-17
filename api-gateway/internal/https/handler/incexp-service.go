package handler

import (
	pb "api-gateway/genproto/incexpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/v1/transactions/income [post]
// @Summary REGISTER INCOME
// @Security  		BearerAuth
// @Description This method registers income
// @Tags TRANSACTIONS
// @Accept json
// @Produce json
// @Param income body models.RegisterIncomeRequest true "Income"
// @Success 200 {object} models.RegisterIncomeResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) RegisterIncome(c *gin.Context) {

	req := pb.RegisterIncomeRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RegisterIncome: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "can't get user id"})
		return
	}
	req.UserId = id.(string)
	
	resp, err := h.Service.RegisterIncome(ctx, &req)
	if err != nil {
		logger.Error("RegisterIncome: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("RegisterIncome: Successfully registered income", resp.TransactionId)
	c.JSON(200, resp)
}

// @Router /api/v1/transactions/expense [post]
// @Summary REGISTER EXPENSE
// @Security  		BearerAuth
// @Description This method registers expense
// @Tags TRANSACTIONS
// @Accept json
// @Produce json
// @Param expense body models.RegisterExpenseRequest true "Expense"
// @Success 200 {object} models.RegisterExpenseResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) RegisterExpense(c *gin.Context) {

	req := pb.RegisterExpenseRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("RegisterExpense: JSON binding error - ", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "can't get user id"})
		return
	}
	req.UserId = id.(string)

	email, exists := c.Get("user_email")
	if !exists {
		c.JSON(500, gin.H{"error": "can't get user email"})
		return
	}
	req.UserEmail = email.(string)

	resp, err := h.Service.RegisterExpense(ctx, &req)
	if err != nil {
		logger.Error("RegisterExpense: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("RegisterExpense: Successfully registered expense", resp.TransactionId)
	c.JSON(200, resp)
}

// @Router /api/v1/transactions [get]
// @Summary GET TRANSACTIONS
// @Security  		BearerAuth
// @Description This method gets income and expense
// @Tags TRANSACTIONS
// @Accept json
// @Produce json
// @Success 200 {object} models.GetListIncomeVSExpenseResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetListIncomeVSExpense(c *gin.Context) {

	req := pb.GetListIncomeVSExpenseRequest{}

	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "can't get user id"})
		return
	}
	req.UserId = id.(string)

	resp, err := h.Service.GetListIncomeVSExpense(ctx, &req)
	if err != nil {
		logger.Error("GetListIncomeVSExpense: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetListIncomeVSExpense: Successfully retrieved transactions - ", len(resp.List))
	c.JSON(200, resp)
}
