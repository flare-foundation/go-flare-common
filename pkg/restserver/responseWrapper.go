package restserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	"gorm.io/gorm"
)

// writeResponse writes a value to the response writer as a JSON object
// Returns an error if the value could not be written
func writeResponse(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&value)
	if err != nil {
		http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
	}
}

func writeResponseError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(map[string]string{"error": message})
	if err != nil {
		http.Error(w, fmt.Sprintf("error writing response: %s", err), http.StatusInternalServerError)
	}
}

// Equivalent to WriteAPIResponse with status APIResponseStatusOk
func WriteAPIResponseOk[T any](w http.ResponseWriter, value T) {
	writeResponse(w, value)
}

// Use for error responses
func WriteAPIResponseError(
	w http.ResponseWriter,
	code int,
	errorMessage string,
) {
	writeResponseError(w, code, errorMessage)
}

// Decode body from the request into value.
// Any error is written into the response and false is returned.
// (It is enough to just return from the request handler on false value)
func DecodeBody(w http.ResponseWriter, r *http.Request, value interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&value)
	if err != nil {
		WriteAPIResponseError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("error parsing request body on endpoint: %s: %s", r.URL.Path, err),
		)
		return false
	}
	err = validate.Struct(value)
	if err != nil {
		WriteAPIResponseError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("error validating request body on endpoint: %s: %s", r.URL.Path, err),
		)
		return false
	}
	return true
}

func DecodeQueryParams(w http.ResponseWriter, r *http.Request, value interface{}) bool {
	decoder := schema.NewDecoder()
	err := decoder.Decode(value, r.URL.Query())
	if err != nil {
		WriteAPIResponseError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("error parsing query params on endpoint: %s", r.URL.Path),
		)
		return false
	}
	err = validate.Struct(value)
	if err != nil {
		WriteAPIResponseError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("error validating query params on endpoint: %s", r.URL.Path),
		)
		return false
	}
	return true
}

func BadParamsErrorHandler(err error) *ErrorHandler {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(
				w,
				http.StatusBadRequest,
				fmt.Sprintf("Error with params: %s", err),
			)
		},
	}
}

func InternalServerErrorHandler(err error) *ErrorHandler {
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(w, http.StatusInternalServerError, "internal server error")
		},
	}
}

func ToEarlyErrorHandler(err error) *ErrorHandler {
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(w, http.StatusBadRequest, "request to early")
		},
	}
}

func NotAvailableErrorHandler(err error) *ErrorHandler {
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(w, http.StatusBadRequest, "data not available yet")
		},
	}
}
