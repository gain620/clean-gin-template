// Package app configures and runs application.
package app

import (
	"clean-gin-template/config"
	v1 "clean-gin-template/internal/controller/http/v1"
	"clean-gin-template/internal/usecase"
	webapi "clean-gin-template/internal/web-api"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/evrone/go-clean-template/pkg/httpserver"
	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	//"github.com/evrone/go-clean-template/config"
	//amqprpc "github.com/evrone/go-clean-template/internal/controller/amqp_rpc"
	//v1 "github.com/evrone/go-clean-template/internal/controller/http/v1"
	//"github.com/evrone/go-clean-template/internal/usecase"
	//"github.com/evrone/go-clean-template/internal/usecase/repo"
	//"github.com/evrone/go-clean-template/internal/usecase/webapi"
	//"github.com/evrone/go-clean-template/pkg/httpserver"
	//"github.com/evrone/go-clean-template/pkg/logger"
	//"github.com/evrone/go-clean-template/pkg/postgres"
	//"github.com/evrone/go-clean-template/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates dependency components for injection.
func Run(cfg *config.Config) {
	l := &log.Logger{
		Out:   os.Stderr,
		Level: log.DebugLevel,
		//Level: cfg.MyLog.Level,
		Formatter: &nested.Formatter{
			HideKeys:    true,
			CallerFirst: true,
			FieldsOrder: []string{"component", "action"},
		},
	}

	// Repository
	//pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	//}
	//defer pg.Close()

	// Use case
	//translationUseCase := usecase.New(
	//	repo.New(pg),
	//	webapi.New(),
	//)
	//
	githubUseCase := usecase.New(
		webapi.New(),
	)

	//// RabbitMQ RPC Server
	//rmqRouter := amqprpc.NewRouter(translationUseCase)
	//
	//rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	//}

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, githubUseCase, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("Close signal detected : " + s.String())
		l.Info("Server quitting... ")
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
		//case err = <-rmqServer.Notify():
		//	l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
