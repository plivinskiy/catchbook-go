package app

import (
	"catchbook/internal/config"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	cfg       *config.Config
	logger    *slog.Logger
	server    *http.Server
	container *Container
}

func NewApp() *Application {
	ctx := context.Background()
	container := NewContainer(ctx)
	return &Application{
		cfg:       container.GetConfig(),
		logger:    container.GetLogger(),
		container: container,
	}
}

func (a *Application) Run() {
	a.logger.Info("start application")
	go func() {
		a.server = &http.Server{
			Addr:    a.cfg.ListenAddress(),
			Handler: a.setupRouter(),
		}
		err := a.server.ListenAndServe()
		if err != nil {
			a.logger.Error(err.Error())
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	s := <-stop
	a.logger.Info("shutdown server ...", slog.Any("signal", s))
	a.Stop()
}

func (a *Application) setupRouter() *gin.Engine {
	gin.SetMode(a.cfg.GinMode)
	r := gin.Default()
	for _, h := range a.container.Handlers() {
		h.Register(r)
	}
	return r
}

func (a *Application) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.GetShutdownTimeout())
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		a.logger.Info("timeout..", slog.Int64("seconds", int64(a.cfg.GetShutdownTimeout())))
	}
	a.logger.Info("Server exiting")
}
