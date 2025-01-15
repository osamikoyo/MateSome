package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/matesome/internal/config"
	"github.com/osamikoyo/matesome/internal/transport/handler"
	"github.com/osamikoyo/matesome/pkg/loger"
	"os"
	"os/signal"
)

type App struct {
	Handler handler.Handler
	logger  loger.Logger
	cfg     config.Config
	ctx     context.Context
	server  *echo.Echo
}

func Init() *App {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	return &App{
		Handler: handler.New(),
		logger:  loger.New(),
		cfg:     config.Get(),
		ctx:     ctx,
		server:  echo.New(),
	}
}

func (a *App) Run() error {
	go func() {
		<-a.ctx.Done()
		a.logger.Info().Msg("Server shutdown!")
		a.server.Shutdown(a.ctx)
	}()

	a.Handler.RegisterRoutes(a.server)

	a.logger.Info().Msg("Http server started on " + fmt.Sprintf("%s:%d :3", a.cfg.Hostname, a.cfg.Port))

	return a.server.Start(fmt.Sprintf("%s:%d", a.cfg.Hostname, a.cfg.Port))
}
