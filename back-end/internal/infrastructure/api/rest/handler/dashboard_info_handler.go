package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *Handler) GetTotalRevenue(c echo.Context) error {
    logger.Debug("Calculating total revenue by storeId...")
    storeId := c.Param("storeId")
    revenue, err := h.dashboardInfoService.GetTotalRevenue(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
	logger.Debug("Got the sales",zap.Float64("sales number ", revenue))
    return c.JSON(http.StatusOK, revenue)
}

func (h *Handler) GetTotalSales(c echo.Context) error {
    logger.Debug("Calculating total sales by storeId...")
    storeId := c.Param("storeId")
    sales, err := h.dashboardInfoService.GetTotalSales(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
	logger.Debug("Got the sales",zap.Int64("sales number ", sales))
    return c.JSON(http.StatusOK, sales)
}


func (h *Handler) GetGraphRevenue(c echo.Context) error {
    logger.Debug("Calculating graph revenue by storeId...")
    storeId := c.Param("storeId")
    graphRevenue, err := h.dashboardInfoService.GetGraphRevenue(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
	logger.Debug("Got the graph revenue",zap.Reflect("graphValue", graphRevenue))
    return c.JSON(http.StatusOK, graphRevenue)
}


