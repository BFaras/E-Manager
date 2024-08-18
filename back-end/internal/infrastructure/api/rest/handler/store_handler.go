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
    userId := c.Param("userId")
    store, err := h.storeRepo.FindByUserId(userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoresByUserId(c echo.Context) error {
    logger.Debug("Fetching all stores by userId...")
    userId := c.Param("userId")
    store, err := h.storeRepo.FindAllByUserId(userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

func (h *Handler) GetStoreByIdAndUserId(c echo.Context) error {
    logger.Debug("Fetching store by id and userId...")
    userId := c.Param("userId")
    id := c.Param("storeId")
    store, err := h.storeRepo.FindByIdAndUserId(id,userId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, store)
}

