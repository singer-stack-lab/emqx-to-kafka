//go:build wireinject
// +build wireinject

package main

import (
	"github.com/singer-stack-lab/emqx-to-kafka/internal/ioc"
	"github.com/singer-stack-lab/emqx-to-kafka/internal/server"

	"github.com/google/wire"
)

func InitBridge() *server.Bridge {
	panic(wire.Build(
		ioc.LoadConfig,
		ioc.InitSarmaClient,
		server.NewBridge,
	))
	return &server.Bridge{}
}
