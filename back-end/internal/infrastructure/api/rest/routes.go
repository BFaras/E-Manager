package rest

import (
	"back-end/internal/infrastructure/api/rest/handler"
	"back-end/internal/infrastructure/api/rest/middleware"
	
) 

func (s *Server) routes(h *handler.Handler, m *middleware.Middleware) {
	s.Echo.Use(m.CORSConfig())
	secured := s.Echo.Group("secured/")
    secured.Use(m.JWTMiddleware)
	
   	s.Echo.GET("users/:userId/store", h.GetStoreByUserId)
	s.Echo.GET("users/:userId/stores", h.GetStoresByUserId)
	s.Echo.GET("users/:userId/stores/:storeId",  h.GetStoreByIdAndUserId)

	s.Echo.GET("stores",  h.GetAllStores)
	s.Echo.GET("stores/:storeId/revenue",  h.GetTotalRevenue)
	s.Echo.GET("stores/:storeId/sales",  h.GetTotalSales)
	s.Echo.GET("stores/:storeId/graphRevenue",  h.GetGraphRevenue)
	s.Echo.GET("stores/:storeId/billboards",  h.GetBillboardsByStoreId)
	s.Echo.GET("stores/:storeId/billboards/active", h.GetActiveBillboardForSpecificStore)
	s.Echo.GET("billboards/:billboardId", h.GetBillboardById)

	secured.DELETE("stores/:storeId/billboards/:billboardId", h.DeleteBillboard)
	secured.POST("stores/:storeId/billboards", h.AddBillboard)
	secured.PATCH("stores/:storeId/billboards/:billboardId", h.UpdateBillboard)

	s.Echo.GET("stores/:storeId/categrories/:categoryId", h.GetCategoryById)
	s.Echo.GET("stores/:storeId/categories", h.GetCategoriesWithBillboard)
	secured.DELETE("stores/:storeId/categories/:categoryId", h.DeleteCategory)
	secured.POST("stores/:storeId/categories", h.AddCategory)
	secured.PATCH("stores/:storeId/categories/:categoryId", h.UpdateCategory)
	
}