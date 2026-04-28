package http_server

import (
	"context"
	"errors"
	"fmt"
	"go-polls-service/internal/core/core_logger"
	"go-polls-service/internal/core_middleware"
	"net/http"

	"go.uber.org/zap"
)

type HTTPServer struct {
	mux        *http.ServeMux
	config     Config
	log        core_logger.Logger
	middleware []core_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, middleware ...core_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:        http.NewServeMux(),
		config:     config,
		log:        *log,
		middleware: middleware,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		s.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)

	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_middleware.ChainMiddleware(s.mux, s.middleware...)

	server := http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)

		s.log.Warn("starting http http_server", zap.String("addr", s.config.Addr))

		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("shutting down http http_server...")

		shutdownCtx, shutdownCancel := context.WithTimeout(
			context.Background(),
			s.config.ShutdownTimeout,
		)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown http http_server: %w", err)
		}

		s.log.Warn("http http_server shutdown completed")
	}

	return nil
}
