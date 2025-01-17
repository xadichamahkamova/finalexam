package handler

import (
	pb "api-gateway/genproto/reportingpb"
	"api-gateway/logger"

	"github.com/gin-gonic/gin"
)

// @Router /api/v1/reports/income-expense [get]
// @Summary GET REPORTS
// @Security  		BearerAuth
// @Description This method gets reports
// @Tags REPORTING
// @Accept json
// @Produce json
// @Success 200 {object} models.GetTotalReportsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetTotalReports(c *gin.Context) {
	
	req := pb.GetTotalReportsRequest{}
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "can't get user id"})
		return
	}
	req.UserId = id.(string)
	
	resp, err := h.Service.GetTotalReports(ctx, &req)
	if err != nil {
		logger.Error("GetTotalReports: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetTotalReports: Successfully retrieved reports")
	c.JSON(200, resp)
}

// @Router /api/v1/reports/spending-by-category [get]
// @Summary GET REPORTS BY CATEGORY
// @Security  		BearerAuth
// @Description This method gets reports by category
// @Tags REPORTING
// @Accept json
// @Produce json
// @Success 200 {object} models.GetReportsByCategoryResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
func (h *HandlerST) GetReportsByCategory(c *gin.Context) {

	req := pb.GetReportsByCategoryRequest{}
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(500, gin.H{"error": "can't get user id"})
		return
	}
	req.UserId = id.(string)
	
	resp, err := h.Service.GetReportsSpendingByCategory(ctx, &req)
	if err != nil {
		logger.Error("GetReportsByCategory: Service error - ", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("GetReportsByCategory: Successfully retrieved reports by category")
	c.JSON(200, resp)
}
