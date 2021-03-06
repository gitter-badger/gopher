package gopher

import "net/http"

type Loggable interface {
	Log() interface{}

	Info(msg string, args ...interface{})

	Debug(msg string, args ...interface{})

	Warn(msg string, args ...interface{})

	Error(msg string, args ...interface{})

	Fatal(msg string, args ...interface{})

	Panic(msg string, args ...interface{})
}

type Parametable interface {
	PathParams(r *http.Request) map[string]string

	PathParam(r *http.Request, param string) string
}

type Renderable interface {
	NewRenderer() interface{}

	View(rw http.ResponseWriter, status int, name string, binding interface{})
}

type Routable interface {
	http.Handler

	Router() interface{}

	Get(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Head(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Post(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Put(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Patch(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Delete(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Options(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Match(verbs []string, path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	All(path string, fn func(http.ResponseWriter, *http.Request)) interface{}

	Serve()
}
