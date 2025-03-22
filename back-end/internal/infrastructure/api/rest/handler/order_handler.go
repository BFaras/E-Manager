package handler

import (
	"back-end/internal/infrastructure/logger"
	"net/http"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"database/sql"
)

type RequestBody struct {
	ProductsId []string `json:"productsId"`
}

type UpdateOrderRequest struct {
	IsPaid  bool   `json:"isPaid"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

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

func (h *Handler) AddOrder(c echo.Context) error {
	storeId := c.Param("storeId")
	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	if len(body.ProductsId) == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "No product IDs provided"})
	}

	order, err := h.orderService.CreateOrder(storeId, body.ProductsId)
	if err != nil {
		logger.Error("Failed to create order", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create order"})
	}

	return c.JSON(http.StatusOK, order)
}

func (h *Handler) UpdateOrder(c echo.Context) error {
    logger.Debug("Updating order after purchase...")

    orderId := c.Param("orderId")

    var req UpdateOrderRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    order, err := h.orderService.GetOrder(orderId)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusNotFound, err.Error())
        }
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    order.IsPaid = req.IsPaid
    order.Address = req.Address
    order.Phone = req.Phone

    if err := h.orderService.UpdateOrder(order); err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    orderItems, err := h.orderItemService.GetOrderItemsByOrderId(order.Id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    for _, orderItem := range orderItems {

        product, err := h.productService.GetProduct(orderItem.ProductId)
        if err != nil {
            logger.Error("Failed to fetch product", zap.String("productId", orderItem.ProductId), zap.Error(err))
            continue 
        }


        product.Count -= 1 
        if product.Count < 1 {
            product.IsArchived = true
        }

        if err := h.productService.UpdateProduct(product); err != nil {
            logger.Error("Failed to update product", zap.String("productId", product.Id), zap.Error(err))
            continue
        }
    }

    return c.NoContent(http.StatusOK)
}