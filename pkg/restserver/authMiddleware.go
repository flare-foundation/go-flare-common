package restserver

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

// ExcludeEndpointStruct identifies a path/method pair to exclude from API key authentication.
type ExcludeEndpointStruct struct {
	PathPattern string
	Method      string
}

// APIKeyAuthMiddleware enforces API key authentication on incoming HTTP requests.
type APIKeyAuthMiddleware struct {
	KeyName          string
	Keys             []string
	ExcludeEndpoints []ExcludeEndpointStruct
	keyMap           map[string]bool
	excludeMap       map[string]map[string]bool
}

// Init initializes internal lookup maps for API keys and excluded endpoints.
func (keyMiddleware *APIKeyAuthMiddleware) Init() {
	keyMiddleware.keyMap = make(map[string]bool)
	keyMiddleware.excludeMap = make(map[string]map[string]bool)
	for _, key := range keyMiddleware.Keys {
		keyMiddleware.keyMap[key] = true
	}
	for _, endpoint := range keyMiddleware.ExcludeEndpoints {
		if keyMiddleware.excludeMap[endpoint.PathPattern] == nil {
			keyMiddleware.excludeMap[endpoint.PathPattern] = make(map[string]bool)
			keyMiddleware.excludeMap[endpoint.PathPattern][endpoint.Method] = true
		} else {
			keyMiddleware.excludeMap[endpoint.PathPattern][endpoint.Method] = true
		}
	}
}

// Middleware returns an HTTP middleware that validates the API key for each request.
func (keyMiddleware *APIKeyAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if path and method combination are in exclude list
		if _, found := keyMiddleware.excludeMap[r.URL.Path][r.Method]; found {
			// Endpoint is excluded from api key check
			next.ServeHTTP(w, r)
			return
		}

		// Check the api key
		token := r.Header.Get(keyMiddleware.KeyName)
		if _, found := keyMiddleware.keyMap[token]; found {
			// API key is valid
			next.ServeHTTP(w, r)
			return
		}

		// Access denied
		errorString := "Unauthorized, provide valid " + keyMiddleware.KeyName + " api key"
		http.Error(w, errorString, http.StatusUnauthorized)
	})
}

// SecuritySchemes returns the OpenAPI security scheme definition for API key authentication.
func (keyMiddleware *APIKeyAuthMiddleware) SecuritySchemes() openapi3.SecuritySchemes {
	return openapi3.SecuritySchemes{
		keyMiddleware.KeyName: &openapi3.SecuritySchemeRef{
			Value: &openapi3.SecurityScheme{
				Type: "apiKey",
				In:   "header",
				Name: keyMiddleware.KeyName,
			},
		},
	}
}
