package gopher

import "net/http"

// Routers

func (c appContainer) Router() Routable {
	return container.router.Router().(Routable)
}

// Parameters

func PathParams(req *http.Request) map[string]string {
	return container.parameters.PathParams(req)
}

func PathParam(req *http.Request, param string) string {
	return container.parameters.PathParam(req, param)
}

// Logger
func (c appContainer) Log() Loggable {
	return container.logger.Log().(Loggable)
}

// Renderer
func NewRenderer() Renderable {
	return container.renderer.NewRenderer().(Renderable)
}

func (c appContainer) View(rw http.ResponseWriter, status int, name string, binding interface{}) {
	container.renderer.View(rw, status, name, binding)
}
