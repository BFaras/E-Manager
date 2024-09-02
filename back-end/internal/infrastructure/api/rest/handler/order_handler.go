package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h* Handler) GetOrderById(c echo.Context) (error) {
	logger.Debug("Fetching order by id...")
    orderId := c.Param("orderId")
    order, err := h.orderService.GetOrder(orderId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, order)
}

func (h *Handler) GetAllOrdersWithExtraInformationByStoreId(c echo.Context) error {
    logger.Debug("Fetch order with orderItem and product...")
    storeId := c.Param("storeId")
    sizes, err := h.orderService.GetAllOrdersWithExtraInformationByStoreId(storeId)
    if err != nil {
        logger.Debug("erreur trouver",zap.Error(err))
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }
    return c.JSON(http.StatusOK, sizes)
}