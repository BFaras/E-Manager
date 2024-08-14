package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetStoreById(c echo.Context) error {
    logger.Debug("Fetching store by storeId...")
    id := c.Param("id")
    store, err := h.storeRepo.FindById(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreByUserId(c echo.Context) error {
    logger.Debug("Fetching store by userId...")
    id := c.Param("userId")
    store, err := h.storeRepo.FindByUserId(id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

