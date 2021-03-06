package service

import (
	"context"

	"github.com/google/wire"
	"github.com/goriller/ginny"
	pb "github.com/goriller/ginny-demo/api/proto"
	"github.com/goriller/ginny-demo/internal/config"
	"github.com/goriller/ginny-demo/internal/repo"
	"github.com/goriller/ginny/server/mux"
)

// ProviderSet
var ProviderSet = wire.NewSet(NewService, RegisterService)

// Service the instance for grpc proto.
type Service struct {
	pb.UnimplementedSayServer
	config *config.Config
	// Introduce new dependencies here, exp:
	userRepository *repo.UserRepo
}

// NewService new service that implement hello
func NewService(
	config *config.Config,
	userRepository *repo.UserRepo,
) *Service {
	mux.RegisterErrorCodes(pb.ErrorCode_name)
	return &Service{
		config:         config,
		userRepository: userRepository,
	}
}

// RegisterService
func RegisterService(ctx context.Context, sev *Service) ginny.RegistrarFunc {
	return func(app *ginny.Application) error {
		// 注册gRPC服务
		app.Server.RegisterService(&pb.Say_ServiceDesc, sev)

		if app.Option.HttpAddr != "" {
			// 注册http服务
			if err := pb.RegisterSayHandlerServer(ctx, app.Server.ServeMux(), sev); err != nil {
				return err
			}
		}
		return nil
	}
}
