package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h* Handler) GetBillboardById(c echo.Context) (error) {
	logger.Debug("Fetching billboard by id...")
    billboardId := c.Param("billboardId")
    billboard, err := h.billboardRepo.FindByID(billboardId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    logger.Debug("got billboard", zap.Reflect("billboard", billboard))
    return c.JSON(http.StatusOK, billboard)
}

func (h *Handler) GetBillboardsByStoreId(c echo.Context) (error) {
	logger.Debug("Fetching all billboards by storeId...")
	storeId := c.Param("storeId")
	store, err := h.billboardRepo.GetBillboardsByStoreId(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	logger.Debug("got all billboard", zap.Reflect("billboards", store))
	return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetActiveBillboardForSpecificStore(c echo.Context) (error) {
	logger.Debug("Fetching active billboard for a specific store...")
	storeId := c.Param("storeId")
	store, err := h.billboardRepo.GetActiveBillboard(storeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	logger.Debug("got all billboard", zap.Reflect("billboards", store))
	return c.JSON(http.StatusOK, store)
}


func (h *Handler) DeleteByBillboardId(c echo.Context) (error) {
	logger.Debug("Deleting billboard by id...")
	billboardId := c.Param("billboardId")
	err := h.billboardRepo.Delete(billboardId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}