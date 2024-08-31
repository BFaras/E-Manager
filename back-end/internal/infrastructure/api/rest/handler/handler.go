package handler

import (
	"back-end/internal/application/service"
)

type Handler struct {
	storeService *service.StoreService
	billboardService *service.BillboardService
	categoryService *service.CategoryService
	colorService *service.ColorService
	imageService *service.ImageService
	orderItemService  *service.OrderItemService
	orderService  *service.OrderService
	productService  *service.ProductService
	sizeService  *service.SizeService
	dashboardInfoService  *service.DashboardInfoService
	
}

func New(storeService *service.StoreService, billboardService *service.BillboardService, categoryService *service.CategoryService,
	colorService *service.ColorService, imageService *service.ImageService, orderItemService *service.OrderItemService, orderService *service.OrderService,
	productService *service.ProductService,sizeService *service.SizeService,dashboardInfoService  *service.DashboardInfoService) *Handler {
	return &Handler{
		storeService: storeService,
		billboardService: billboardService,
		categoryService: categoryService,
		colorService: colorService,
        imageService: imageService,
        orderItemService: orderItemService,
        orderService: orderService,
        productService: productService,
		sizeService: sizeService,
		dashboardInfoService: dashboardInfoService,
	}
}
