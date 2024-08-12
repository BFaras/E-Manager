package handler

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func (h *Handler) GetStore(c echo.Context) {
    id := c.Param("id")
    store, err := h.storeRepo.FindByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, store)
}
