package rest

import (
	"back-end/internal/infrastructure/api/rest/handler"
	"back-end/internal/infrastructure/api/rest/middleware"
	"back-end/internal/infrastructure/api/rest/validator"
	"back-end/internal/infrastructure/db"
	"context"
	"database/sql"
	"net"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/juju/errors"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


 type Server struct {
	echo *echo.Echo
	log  *zap.Logger
	cfg  *viper.Viper
	db   *sql.DB
 }

 func New(cfg *viper.Viper, log *zap.Logger) (*Server, error) {
	/*create fb in here and use .env to connect*/

	server := &Server{
	   echo: echo.New(),
	   log:  log,
	   cfg:  cfg,
	}

	server.log.Info("Start setting up the server...")

	var err error

	server.db, err = db.SetUpDatabase(server)

    if err != nil {
        return nil, err
    }

	server.configure(cfg.Sub("setting"))
	
	storeRepo := db.NewStoreRepository(server.db)

	server.routes(
	   handler.New(storeRepo),
	   middleware.New(),
	)
	server.log.Debug("Successfully connected the  handlers and middlewares to the server")
 
	return server, nil
 }
 
 func (s *Server) Start(ctx context.Context) error {
	errorChan := make(chan error, 1)
 
	go s.start(errorChan)
 
	select {
	case <-ctx.Done():
	   s.log.Info("Shutting down the server")
	   if shutdownErr := s.echo.Shutdown(ctx); shutdownErr != nil {
		  s.log.Error("Error shutting down the server", zap.Error(shutdownErr))
		  return shutdownErr
	   }
	case err := <-errorChan:
	   s.log.Fatal("Failed to start HTTP server", zap.Error(err))
	   return err
	}
 
	return nil
 }
 
 func (s *Server) start(errorChan chan<- error) {
	defer close(errorChan)
 
 
	if err := s.echo.Start(
	   net.JoinHostPort(
		  s.cfg.GetString("host"),
		  s.cfg.GetString("port"),
	   ),
	); err != nil && !errors.Is(err, http.ErrServerClosed) {
	   errorChan <- err
	}
 }
 
 func (s *Server) configure(cfg *viper.Viper) {
	if cfg.IsSet("debug") {
	   s.echo.Debug = cfg.GetBool("debug")
	}
 
	if cfg.IsSet("hide_banner") {
	   s.echo.HideBanner = cfg.GetBool("hide_banner")
	}
 
	if cfg.IsSet("hide_port") {
	   s.echo.HidePort = cfg.GetBool("hide_port")
	}
 
	s.echo.Validator = validator.New()
	s.echo.HTTPErrorHandler = handleErrors(s.log, cfg.GetBool("debug"))
 }
 
 func handleErrors(log *zap.Logger, debug bool) echo.HTTPErrorHandler {
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
			 log.Error("An error occurred", zap.Error(err))
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