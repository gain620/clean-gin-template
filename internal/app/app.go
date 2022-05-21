// Package app configures and runs application.
package app

import (
	"clean-gin-template/config"
	v1 "clean-gin-template/internal/controller/http/v1"
	"clean-gin-template/internal/usecase"
	webapi "clean-gin-template/internal/web-api"
	"clean-gin-template/pkg/db"
	"clean-gin-template/pkg/logger"
	"clean-gin-template/pkg/server"
	"fmt"
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
	l := logger.LogurusSetup(cfg)

	// Repository
	//pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	//}
	//defer pg.Close()
	client, err := db.GetClient(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - db.GetClient: %v", err))
	}
	defer client.Close()

	err = client.Ping()
	if err != nil {
		l.Fatal(fmt.Errorf("db connection error : %v", err))
	}

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

	// Server init
	handler := gin.New()

	// Dependency Injection
	v1.NewRouter(handler, githubUseCase, l)

	// Start http server
	httpServer := server.New(handler, server.Port(cfg.Server.Port))
	//httpsServer := server.New(handler, server.TLS(cfg.Server.Cert, cfg.Server.Key), server.Port(cfg.Server.Port))

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
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	//err = rmqServer.Shutdown()
	//if err != nil {
	//	l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	//}
}
