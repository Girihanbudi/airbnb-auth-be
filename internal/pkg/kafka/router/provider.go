package router

import (
	"airbnb-auth-be/internal/pkg/kafka/router/config"
	"strings"
)

const Instance string = "Kafka Router"

type Options struct {
	config.Config
}

type Router struct {
	Options
	basePath string
	Handlers []Handler
}

func NewRouter(options Options) *Router {
	// set default separator
	if options.Separator == "" {
		options.Separator = "."
	}

	return &Router{
		Options: options,
	}
}

func (r *Router) Group(relativePath string) *Router {
	r.basePath = r.calculateAbsolutePath(relativePath)
	return r
}

func (r *Router) Listen(relativePath string, handler EventHandler) {
	absolutePath := r.calculateAbsolutePath(relativePath)
	r.Handlers = append(r.Handlers, Handler{
		Topic:   absolutePath,
		Handler: handler,
	})
}

func (r *Router) calculateAbsolutePath(relativePath string) string {
	return r.joinPaths(r.basePath, relativePath)
}

func (r *Router) joinPaths(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}

	if absolutePath == "" {
		return relativePath
	}

	return strings.Join([]string{absolutePath, relativePath}, r.Separator)
}
