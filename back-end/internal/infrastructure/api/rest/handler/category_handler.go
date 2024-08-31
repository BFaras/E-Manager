package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)



func (h *Handler) GetCategoriesWithBillboard(c echo.Context) error {
    storeId := c.Param("storeId")

    categories, err := h.categoryService.GetCategoriesWithBillboard(storeId)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, zap.Error(err))
    }

    return c.JSON(http.StatusOK, categories)
}

