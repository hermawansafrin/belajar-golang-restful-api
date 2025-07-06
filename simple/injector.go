//go:build wireinject
// +build wireinject

// isinya nanti DI untuk autogenerate google wire
package simple

import "github.com/google/wire"

func InitializedService() *SimpleService {
	// fungsi mana yang akan dibuatkan DI
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil // nanti akan di ganti sama google wire
}
