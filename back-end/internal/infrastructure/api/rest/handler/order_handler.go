package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTotalRevenue(c echo.Context) error {
    logger.Debug("Fetching total revenue by storeId...")
    storeId := c.Param("storeId")
    revenue, err := h.orderRepo.CalculateRevenue(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, revenue)
}

func (h *Handler) GetTotalSales(c echo.Context) error {
    logger.Debug("Fetching total sales by storeId...")
    storeId := c.Param("storeId")
    sales, err := h.orderRepo.CalculateSales(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, sales)
}

