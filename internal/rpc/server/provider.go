package server

import (
	"github.com/gorillazer/ginny-demo/api/proto"

	"github.com/google/wire"
	"github.com/gorillazer/ginny-serve/grpc"
	stdGrpc "google.golang.org/grpc"
)

// CreateInitServerFn
func CreateInitServerFn(
	d *DetailsServer,
) grpc.InitServers {
	return func(s *stdGrpc.Server) {
		proto.RegisterDetailsServer(s, d)
	}
}

// ProviderSet
var ProviderSet = wire.NewSet(

	NewDetailsServer,
	CreateInitServerFn,
)
