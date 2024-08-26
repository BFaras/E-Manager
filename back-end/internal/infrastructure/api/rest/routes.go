package rest

import (
	"back-end/internal/infrastructure/api/rest/handler"
	"back-end/internal/infrastructure/api/rest/middleware"
	
) 

func (s *Server) routes(h *handler.Handler, m *middleware.Middleware) {
	s.Echo.Use(m.CORSConfig())
   	s.Echo.GET("/users/:userId/store", h.GetStoreByUserId)
	s.Echo.GET("/users/:userId/stores", h.GetStoresByUserId)
	s.Echo.GET("/users/:userId/stores/:storeId",  h.GetStoreByIdAndUserId)
	
	/*
	s.Echo.DELETE("/users/:userId/stores/:storeId/billboards/:billboardId",  h.GetBillboardsByStoreId)
	*/

	s.Echo.GET("stores/:storeId/revenue",  h.GetTotalRevenue)
	s.Echo.GET("stores/:storeId/sales",  h.GetTotalSales)
	s.Echo.GET("stores/:storeId/graphRevenue",  h.GetGraphRevenue)
	s.Echo.GET("stores/:storeId/billboards",  h.GetBillboardsByStoreId)
	s.Echo.GET("stores/:storeId/billboards/active", h.GetActiveBillboardForSpecificStore)

	
}