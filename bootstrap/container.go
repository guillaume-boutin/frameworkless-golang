package bootstrap

import (
	"github.com/guillaume-boutin/frameworkless-golang/app/provider"
	"github.com/sarulabs/di"
)

func registerProviders(builder *di.Builder) *di.Builder {

	builder.Add(provider.TestProvider)

	return builder
}

func Container() di.Container {
	builder, _ := di.NewBuilder()

	builder = registerProviders(builder)

	return builder.Build()
}
