package rest

import (
	"back-end/internal/infrastructure/api/rest/handler"
	"back-end/internal/infrastructure/api/rest/middleware"
	"back-end/internal/infrastructure/api/rest/validator"
	"back-end/internal/infrastructure/db"
	"back-end/internal/infrastructure/logger"

	"context"
	"database/sql"
	"net"
	"net/http"

	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


 type Server struct {
	Echo *echo.Echo
	Cfg  *viper.Viper
	Db   *sql.DB
 }

 func New(cfg *viper.Viper) (*Server, error) {
	/*create fb in here and use .env to connect*/

	server := &Server{
	   Echo: echo.New(),
	   Cfg:  cfg,
	}

	logger.Info("Start setting up the server...")

	var err error


	server.Db, err = db.SetUpDatabase()

    if err != nil {
        return nil, err
    }

	server.configure(cfg.Sub("setting"))
	
	storeRepo := db.NewStoreRepository(server.Db)

	server.routes(
	   handler.New(storeRepo),
	   middleware.New(),
	)
	logger.Debug("Successfully connected the  handlers and middlewares to the server")
 
	return server, nil
 }
 
 func (s *Server) Start(ctx context.Context) error {
	errorChan := make(chan error, 1)
 
	go s.start(errorChan)
 
	select {
	case <-ctx.Done():
		logger.Info("Shutting down the server")
	   if shutdownErr := s.Echo.Shutdown(ctx); shutdownErr != nil {
		  logger.Error("Error shutting down the server", zap.Error(shutdownErr))
		  return shutdownErr
	   }
	case err := <-errorChan:
	   logger.Fatal("Failed to start HTTP server", zap.Error(err))
	   return err
	}
 
	return nil
 }
 
 func (s *Server) start(errorChan chan<- error) {
	defer close(errorChan)
 
 
	if err := s.Echo.Start(
	   net.JoinHostPort(
		  s.Cfg.GetString("host"),
		  s.Cfg.GetString("port"),
	   ),
	); err != nil && !errors.Is(err, http.ErrServerClosed) {
	   errorChan <- err
	}
 }
 
 func (s *Server) configure(cfg *viper.Viper) {
	if cfg.IsSet("debug") {
	   s.Echo.Debug = cfg.GetBool("debug")
	}
 
	if cfg.IsSet("hide_banner") {
	   s.Echo.HideBanner = cfg.GetBool("hide_banner")
	}
 
	if cfg.IsSet("hide_port") {
	   s.Echo.HidePort = cfg.GetBool("hide_port")
	}
 
	s.Echo.Validator = validator.New()
	s.Echo.HTTPErrorHandler = handleErrors(cfg.GetBool("debug"))
 }
 
 func handleErrors(debug bool) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
	   var (
		  code       = http.StatusInternalServerError
		  msg        string
		  errorStack any
	   )
 
	   if he, ok := err.(*echo.HTTPError); ok {
		  code = he.Code
		  msg = he.Message.(string)
	   } else {
		  msg = err.Error()
		  switch true {
		  case errors.Is(err, errors.BadRequest):
			 code = http.StatusBadRequest
		  case errors.Is(err, errors.Forbidden):
			 code = http.StatusForbidden
		  case errors.Is(err, errors.Unauthorized):
			 code = http.StatusUnauthorized
		  case errors.Is(err, errors.NotFound):
			 code = http.StatusNotFound
		  case errors.Is(err, errors.AlreadyExists):
			 code = http.StatusConflict
		  }
 
		  if debug {
			 errorStack = errors.ErrorStack(err)
		  }
	   }
 
	   if !c.Response().Committed {
		  if err != nil && code == http.StatusInternalServerError {
			logger.Error("An error occurred", zap.Error(err))
		  }
 
		  if c.Request().Method == echo.HEAD {
			 err = c.NoContent(code)
		  } else {
			 m := echo.Map{"error": msg}
			 if errorStack != nil {
				m["errorStack"] = errorStack
			 }
			 err = c.JSON(code, m)
		  }
	   }
	}
 }