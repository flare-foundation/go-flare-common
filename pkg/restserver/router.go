// Package restserver implements a rest server framework that allow developers to create REST APIs with optional Swagger documentation support.
//
// Deprecated: This package is deprecated and no longer maintained. Please use alternative solutions.
package restserver

import (
	swagger "github.com/davidebianchi/gswagger"
	"github.com/davidebianchi/gswagger/support/gorilla"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/swaggest/swgui"
	v3 "github.com/swaggest/swgui/v3"
)

// Router defines the interface for registering routes, middleware, and sub-routers.
type Router interface {
	AddRoute(path string, handler RouteHandler, description ...string)
	AddMiddleware(middleware mux.MiddlewareFunc)
	WithPrefix(prefix string, tag string) Router
	Finalize()
}

////////////////////////////////////////////////////////
//////// DEFAULT ROUTER IMPLEMENTATION /////////////////
////////////////////////////////////////////////////////

// Default router implementation using gorilla/mux.
type defaultRouter struct {
	router *mux.Router
}

func (r *defaultRouter) AddRoute(path string, handler RouteHandler, description ...string) {
	r.router.HandleFunc(path, handler.Handler).Methods(handler.Method)
}

func (r *defaultRouter) AddMiddleware(middleware mux.MiddlewareFunc) {
	r.router.Use(middleware)
}

func (r *defaultRouter) WithPrefix(prefix string, tag string) Router {
	return &defaultRouter{
		router: r.router.PathPrefix(prefix).Subrouter(),
	}
}

func (r *defaultRouter) Finalize() {
}

// NewDefaultRouter creates a Router backed by the given gorilla/mux router.
func NewDefaultRouter(mRouter *mux.Router) Router {
	return &defaultRouter{
		router: mRouter,
	}
}

////////////////////////////////////////////////////////
//////// SWAGGER ROUTER IMPLEMENTATION /////////////////
////////////////////////////////////////////////////////

// SwaggerRouterConfig holds configuration for a swagger-enabled router.
type SwaggerRouterConfig struct {
	Title           string
	Version         string
	SwaggerBasePath string
	SecuritySchemes openapi3.SecuritySchemes
}

// JSONDocumentationPath returns the path for the JSON OpenAPI documentation endpoint.
func (c *SwaggerRouterConfig) JSONDocumentationPath() string {
	return c.SwaggerBasePath + "-json"
}

// YAMLDocumentationPath returns the path for the YAML OpenAPI documentation endpoint.
func (c *SwaggerRouterConfig) YAMLDocumentationPath() string {
	return c.SwaggerBasePath + "-yaml"
}

// Router implementation with swagger support.
type swaggerRouter struct {
	mRouter *mux.Router
	router  *swagger.Router[gorilla.HandlerFunc, *mux.Route]
	tag     string
	config  SwaggerRouterConfig
}

// NewSwaggerRouter creates a Router with OpenAPI documentation support.
func NewSwaggerRouter(mRouter *mux.Router, config SwaggerRouterConfig) Router {
	router, _ := swagger.NewRouter(gorilla.NewRouter(mRouter), swagger.Options{
		Openapi: &openapi3.T{
			Info: &openapi3.Info{
				Title:   config.Title,
				Version: config.Version,
			},
			// Security: secReq,
			Components: &openapi3.Components{
				SecuritySchemes: config.SecuritySchemes,
			},
		},
		JSONDocumentationPath: config.JSONDocumentationPath(),
		YAMLDocumentationPath: config.YAMLDocumentationPath(),
	})
	return &swaggerRouter{
		mRouter: mRouter,
		router:  router,
		tag:     "",
		config:  config,
	}
}

// AddRoute adds a route and generates OpenAPI definitions from the handler.
// The first description item sets the summary field; the second sets the description field.
func (r *swaggerRouter) AddRoute(path string, handler RouteHandler, description ...string) {
	swaggerDefinitions := handler.SwaggerDefinitions
	swaggerDefinitions.Tags = []string{r.tag}
	if len(description) > 0 {
		swaggerDefinitions.Summary = description[0]
		if len(description) > 1 {
			swaggerDefinitions.Description = description[1]
		}
	}

	_, err := r.router.AddRoute(handler.Method, path, handler.Handler, swaggerDefinitions)
	if err != nil {
		logger.Fatal(err)
	}
}

func (r *swaggerRouter) AddMiddleware(middleware mux.MiddlewareFunc) {
	r.mRouter.Use(middleware)
}

func (r *swaggerRouter) WithPrefix(prefix string, tag string) Router {
	mSubRouter := r.mRouter.NewRoute().Subrouter()
	subRouter, err := r.router.SubRouter(gorilla.NewRouter(mSubRouter), swagger.SubRouterOptions{
		PathPrefix: prefix,
	})
	if err != nil {
		logger.Panic(err)
	}

	return &swaggerRouter{
		mRouter: mSubRouter,
		router:  subRouter,
		tag:     tag,
	}
}

func (r *swaggerRouter) Finalize() {
	if err := r.router.GenerateAndExposeOpenapi(); err != nil {
		logger.Fatal(err)
	}

	config := swgui.Config{
		Title:       r.config.Title,
		SwaggerJSON: r.config.JSONDocumentationPath(),
		BasePath:    r.config.SwaggerBasePath,
		HideCurl:    false,
		ShowTopBar:  false,
	}

	handler := v3.NewHandlerWithConfig(config)
	r.mRouter.PathPrefix(r.config.SwaggerBasePath).HandlerFunc(handler.ServeHTTP)
}
