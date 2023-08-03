package service

import (
	"github.com/google/wire"
	"service/internal/boundedcontexts/hello"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(hello.InjectHelloRepository, hello.InjectHelloGrpcHandler, NewHelloServer)
