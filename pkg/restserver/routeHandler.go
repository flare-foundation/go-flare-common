package restserver

import (
	"net/http"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/gorilla/mux"
)

// RouteHandler combines an HTTP handler with its OpenAPI definitions and HTTP method.
type RouteHandler struct {
	Handler            func(w http.ResponseWriter, r *http.Request)
	SwaggerDefinitions swagger.Definitions
	Method             string // Take from net/http package (MethodGet, MethodPost, etc)
}

// ErrorHandler wraps a function that writes an error response.
type ErrorHandler struct {
	Handler func(w http.ResponseWriter)
}

// GeneralRouteHandler creates a RouteHandler that parses path parameters, query parameters of type Q,
// and a request body of type B, then delegates to the handler function and returns a JSON response of type R.
// OpenAPI definitions are generated from paramDescriptions, queryObject, bodyObject, and respObject.
func GeneralRouteHandler[Q any, B any, R any](
	handler func(map[string]string, Q, B) (R, *ErrorHandler),
	method string,
	responseCode int,
	paramDescriptions map[string]string, // Path params descriptions for openAPI
	queryObject Q,
	bodyObject B,
	respObject R,
	security []string,
) RouteHandler {
	routeHandler := func(w http.ResponseWriter, r *http.Request) {
		var body B
		if !IsNil(bodyObject) && !DecodeBody(w, r, &body) {
			return
		}

		var query Q
		if !IsNil(queryObject) && !DecodeQueryParams(w, r, &query) {
			return
		}
		params := mux.Vars(r)

		resp, err := handler(params, query, body)
		if err != nil {
			err.Handler(w)
			return
		}
		WriteAPIResponseOk(w, resp)
	}

	// Swagger definitions
	pathParams := createPathParamsDescription(paramDescriptions)
	queryString := createQueryDescription(queryObject)
	requestBody := createRequestBodyDescription(bodyObject)
	sec := createSecuritiesArray(security)

	swaggerDefinitions := swagger.Definitions{
		RequestBody: requestBody,
		PathParams:  pathParams,
		Querystring: queryString,
		Responses: map[int]swagger.ContentValue{
			responseCode: {
				Content: swagger.Content{
					"application/json": {Value: respObject},
				},
			},
		},
		Security: sec,
	}
	return RouteHandler{
		Handler:            routeHandler,
		SwaggerDefinitions: swaggerDefinitions,
		Method:             method,
	}
}

// Create a security object for openAPI from a list of security names.
func createSecuritiesArray(security []string) swagger.SecurityRequirements {
	if security == nil {
		return nil
	}
	sec := make(swagger.SecurityRequirement)
	for _, element := range security {
		sec[element] = []string{}
	}
	ret := make(swagger.SecurityRequirements, 0)
	ret = append(ret, sec)
	return ret
}

// Create openAPI path parameters description from a map of parameter names and descriptions.
func createPathParamsDescription(paramDescriptions map[string]string) map[string]swagger.Parameter {
	if len(paramDescriptions) == 0 {
		return nil
	}

	pathParams := make(map[string]swagger.Parameter)
	for name, description := range paramDescriptions {
		pathParams[name] = swagger.Parameter{
			Schema:      &swagger.Schema{Value: ""},
			Description: description,
		}
	}
	return pathParams
}

// Create openAPI query parameters description from a struct.
func createQueryDescription(queryObject any) swagger.ParameterValue {
	if queryObject == nil {
		return nil
	}
	fields := StructFields(queryObject)
	if len(fields) == 0 {
		return nil
	}

	queryString := make(swagger.ParameterValue)
	for _, field := range fields {
		name := field.Tag.Get("json")
		if name == "" {
			name = field.Name
		}
		queryString[name] = swagger.Parameter{
			Schema:      &swagger.Schema{Value: ""},
			Description: field.Tag.Get("jsonschema"),
		}
	}
	return queryString
}

// Create openAPI request body description from a struct.
func createRequestBodyDescription(bodyObject any) *swagger.ContentValue {
	if bodyObject == nil {
		return nil
	}
	return &swagger.ContentValue{
		Content: swagger.Content{
			"application/json": {Value: bodyObject},
		},
	}
}
