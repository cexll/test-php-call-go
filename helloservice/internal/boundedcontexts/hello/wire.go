package hello

import (
	"service/internal/boundedcontexts/hello/application/handlers"
	"service/internal/boundedcontexts/hello/domain/repositories"
	repositories2 "service/internal/boundedcontexts/hello/infrastructure/repositories"
	"service/internal/svc"
)

func InjectHelloRepository(svc *svc.ServiceContext) repositories.IHelloRepository {
	return &repositories2.HelloRepository{Svc: svc}
}
func InjectHelloGrpcHandler(helloRepo repositories.IHelloRepository) *handlers.HelloGrpcHandler {
	return handlers.NewHelloGrpcHandler(helloRepo)
}
