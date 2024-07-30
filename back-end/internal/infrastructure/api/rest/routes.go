package rest

import (
	"back-end/internal/infrastructure/api/rest/handler"
	"back-end/internal/infrastructure/api/rest/middleware"
) 

func (s *Server) routes(h *handler.Handler, m *middleware.Middleware) {
   s.echo.GET("/hello", h.HelloWold)
}