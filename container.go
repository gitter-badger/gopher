package gopher

import "github.com/gopherlabs/gopher/providers"

var container = appContainer{}

const (
	PROVIDER_LOGGER = iota
	PROVIDER_ROUTER
	PROVIDER_RENDERER
	PROVIDER_PARAMS
)

type appContainer struct {
	logger     Loggable
	router     Routable
	parameters Parametable
	renderer   Renderable
}

func App() *appContainer {
	registerProviders()
	container.Log().Info("Starting Gopher...")
	return &container
}

func registerProviders() {
	RegisterProvider(PROVIDER_LOGGER, providers.LogProvider{})
	RegisterProvider(PROVIDER_ROUTER, providers.RouteProvider{})
	RegisterProvider(PROVIDER_RENDERER, providers.RenderProvider{})
	RegisterProvider(PROVIDER_PARAMS, providers.ParameterProvider{})
}

func RegisterProvider(providerConst int, provider interface{}) {
	switch providerConst {
	case PROVIDER_LOGGER:
		container.logger = provider.(Loggable)
	case PROVIDER_ROUTER:
		container.router = provider.(Routable)
	case PROVIDER_RENDERER:
		container.renderer = provider.(Renderable)
	case PROVIDER_PARAMS:
		container.parameters = provider.(Parametable)
	}
}
