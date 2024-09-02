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

	secured.DELETE("stores/:storeId", h.DeleteStore)
	secured.PATCH("stores/:storeId", h.UpdateStore)
	secured.POST("stores", h.AddStore)

	secured.DELETE("stores/:storeId/billboards/:billboardId", h.DeleteBillboard)
	secured.POST("stores/:storeId/billboards", h.AddBillboard)
	secured.PATCH("stores/:storeId/billboards/:billboardId", h.UpdateBillboard)

	s.Echo.GET("stores/:storeId/categrories/:categoryId", h.GetCategoryById)
	s.Echo.GET("stores/:storeId/categories", h.GetCategoriesWithBillboard)
	secured.DELETE("stores/:storeId/categories/:categoryId", h.DeleteCategory)
	secured.POST("stores/:storeId/categories", h.AddCategory)
	secured.PATCH("stores/:storeId/categories/:categoryId", h.UpdateCategory)

	s.Echo.GET("stores/:storeId/sizes/:sizeId", h.GetSizeById)
	s.Echo.GET("stores/:storeId/sizes", h.GetAllSizes)
	secured.DELETE("stores/:storeId/sizes/:sizeId", h.DeleteSize)
	secured.POST("stores/:storeId/sizes", h.AddSize)
	secured.PATCH("stores/:storeId/sizes/:sizeId", h.UpdateSize)

	s.Echo.GET("stores/:storeId/colors/:colorId", h.GetColorById)
	s.Echo.GET("stores/:storeId/colors", h.GetAllColors)
	secured.DELETE("stores/:storeId/colors/:colorId", h.DeleteColor)
	secured.POST("stores/:storeId/colors", h.AddColor)
	secured.PATCH("stores/:storeId/colors/:colorId", h.UpdateColor)

	s.Echo.GET("stores/:storeId/products/:productId", h.GetProductById)
	s.Echo.GET("stores/:storeId/products", h.GetAllProductsWithExtraInformationByStoreId)
	s.Echo.GET("stores/:storeId/products/:productId/image", h.GetAllProductsWithImageById)
	secured.DELETE("stores/:storeId/products/:productId", h.DeleteProduct)
	secured.POST("stores/:storeId/products", h.AddProduct)
	secured.PATCH("stores/:storeId/products/:productId", h.UpdateProduct)

	s.Echo.GET("stores/:storeId/orders", h.GetAllOrdersWithExtraInformationByStoreId)

}