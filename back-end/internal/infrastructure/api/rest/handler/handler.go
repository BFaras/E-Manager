package handler

import "back-end/internal/domain/repository"

type Handler struct {
	storeRepo repository.StoreRepository 
	
}

func New(storeRepo repository.StoreRepository) *Handler {
	return &Handler{
		storeRepo: storeRepo,
	}
}
