//go:build wireinject
// +build wireinject

package main

import (
	"github.com/goriller/ginny-demo/internal/config"
	"github.com/goriller/ginny-demo/internal/logic"
	"github.com/goriller/ginny-demo/internal/repo"
	"github.com/goriller/ginny-demo/internal/service"

	"github.com/google/wire"
	"github.com/goriller/ginny"
	"github.com/goriller/ginny/server"
	// consul "github.com/goriller/ginny-consul"
	// jaeger "github.com/goriller/ginny-jaeger"
	// "github.com/opentracing/opentracing-go"
)

// NewApp
func NewApp() (*ginny.Application, error) {
	panic(wire.Build(wire.NewSet(
		// consul.ProviderSet,
		// jaeger.ProviderSet,
		config.ProviderSet,
		repo.ProviderSet,
		logic.ProviderSet,
		service.ProviderSet,
		serverOption,
		ginny.AppProviderSet,
	)))
}

func serverOption(
// consul *consul.Client,
// tracer opentracing.Tracer,
) (opts []server.Option) {
	// opts = append(opts, server.WithDiscover(consul))
	// opts = append(opts, server.WithTracer(tracer))
	return
}
