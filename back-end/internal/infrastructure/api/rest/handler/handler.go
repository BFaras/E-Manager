package handler

import "back-end/internal/domain/repository"

type Handler struct {
	storeRepo repository.StoreRepository
	billboardRepo repository.BillboardRepository 
	categoryRepo repository.CategoryRepository
	colorRepo repository.ColorRepository
	imageRepo repository.ImageRepository
	orderItemRepo  repository.OrderItemRepository
	orderRepo  repository.OrderRepository
	productRepo  repository.ProductRepository
	sizeRepo  repository.SizeRepository


	
}

func New(storeRepo repository.StoreRepository, billboardRepo repository.BillboardRepository, categoryRepo repository.CategoryRepository,
	colorRepo repository.ColorRepository, imageRepo repository.ImageRepository, orderItemRepo repository.OrderItemRepository, orderRepo repository.OrderRepository,
	productRepo repository.ProductRepository,sizeRepo repository.SizeRepository) *Handler {
	return &Handler{
		storeRepo: storeRepo,
		billboardRepo: billboardRepo,
		categoryRepo: categoryRepo,
		colorRepo: colorRepo,
        imageRepo: imageRepo,
        orderItemRepo: orderItemRepo,
        orderRepo: orderRepo,
        productRepo: productRepo,
		sizeRepo: sizeRepo,
	}
}
