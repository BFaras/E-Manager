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
}