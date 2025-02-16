package restserver

import (
	"net/http"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/gorilla/mux"
)

type RouteHandler struct {
	Handler            func(w http.ResponseWriter, r *http.Request)
	SwaggerDefinitions swagger.Definitions
	Method             string // Take from net/http package (MethodGet, MethodPost, etc)
}

type ErrorHandler struct {
	Handler func(w http.ResponseWriter)
}

// Route handler factory
// The value passed to handler are the path parameters parsed to a map of string, the query parameters parsed to a struct
// of type Q and the request body parsed to a struct of type B. The response of handler is wrapped to an
// APIResponseWrapper object and returned as json. OpenAPI definitions for the path parameters are generated from the
// paramDescriptions map, definitions for the query parameters are generated from the queryObject and definitions for the
// request body are generated from the bodyObject.
func GeneralRouteHandler[Q interface{}, B interface{}, R interface{}](
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

// Create a security object for openAPI from a list of security names
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

// Create openAPI path parameters description from a map of parameter names and descriptions
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

// Create openAPI query parameters description from a struct
func createQueryDescription(queryObject interface{}) swagger.ParameterValue {
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

// Create openAPI request body description from a struct
func createRequestBodyDescription(bodyObject interface{}) *swagger.ContentValue {
	if bodyObject == nil {
		return nil
	}
	return &swagger.ContentValue{
		Content: swagger.Content{
			"application/json": {Value: bodyObject},
		},
	}
}
