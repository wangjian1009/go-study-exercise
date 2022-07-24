//go:build wireinject
// +build wireinject

package main

import (
	"projectexample/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wirel"
)

// wireApp init kratos application.
func wireApp(*conf.Conf, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz, service.ProviderSet, newApp))
}
