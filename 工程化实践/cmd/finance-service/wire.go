//go:build wireinject
// +build wireinject

package main

import (
	"projectexample/internal/biz/"
	"projectexample/internal/conf"
	"projectexample/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// wireApp init kratos application.
func wireApp(*conf.Conf, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.NewFinanceMgr, service.ProviderSet, newApp))
}
