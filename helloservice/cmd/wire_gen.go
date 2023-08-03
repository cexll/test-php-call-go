// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"service/internal/boundedcontexts/hello"
	"service/internal/config"
	"service/internal/db"
	"service/internal/server"
	"service/internal/service"
	"service/internal/svc"
)

// Injectors from wire.go:

// initApp init app application.
func initApp(c *config.Config) (*server.AppServer, error) {
	serviceContext := svc.NewServiceContext(c)
	iHelloRepository := hello.InjectHelloRepository(serviceContext)
	helloGrpcHandler := hello.InjectHelloGrpcHandler(iHelloRepository)
	helloService := service.NewHelloServer(serviceContext, helloGrpcHandler)
	httpServer := server.NewHttpServer(c, helloService)
	rpcServer := server.NewGrpcServer(c, serviceContext, helloGrpcHandler)
	appServer, err := server.NewApp(serviceContext, helloService, httpServer, rpcServer)
	if err != nil {
		return nil, err
	}
	return appServer, nil
}

func wireMigrate(c *config.Config) *db.Migrator {
	v := db.CreateDataMigrations()
	serviceContext := svc.NewServiceContext(c)
	migrator := db.NewMigrator(v, serviceContext)
	return migrator
}