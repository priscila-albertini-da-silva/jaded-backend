package serverfx

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type params struct {
	fx.In
	Rt []Route
}

func lifecycle(lf fx.Lifecycle, p params) {
	router := gin.Default()

	api := router.Group("/")

	for _, r := range p.Rt {
		handlers := []gin.HandlerFunc{r.Handler}
		api.Handle(r.Method, r.Path, handlers...)
	}

	addr := net.JoinHostPort("host", "port")

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {

				log.Infof("Starting Server at %s...", time.Now())
				log.Infof("Listening on %v...\n", addr)

				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping Server...")

			err := server.Shutdown(ctx)
			if err != nil {
				log.Errorf("Server stop failed: %+v", err)
			}
			return err
		},
	})
}

var ModuleServer = fx.Module("server", fx.Invoke(lifecycle))
