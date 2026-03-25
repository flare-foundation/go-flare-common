package restserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"

	"gorm.io/gorm"
)

// writeResponse writes a value to the response writer as a JSON object.
// Returns an error if the value could not be written.
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

// WriteAPIResponseOk writes a JSON response with the given value.
func WriteAPIResponseOk[T any](w http.ResponseWriter, value T) {
	writeResponse(w, value)
}

// WriteAPIResponseError writes a JSON error response with the given status code and message.
func WriteAPIResponseError(
	w http.ResponseWriter,
	code int,
	errorMessage string,
) {
	writeResponseError(w, code, errorMessage)
}

// DecodeBody decodes and validates the JSON request body into value.
// On error, it writes an error response and returns false.
func DecodeBody(w http.ResponseWriter, r *http.Request, value any) bool {
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

// DecodeQueryParams decodes and validates URL query parameters into value.
// On error, it writes an error response and returns false.
func DecodeQueryParams(w http.ResponseWriter, r *http.Request, value any) bool {
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

// BadParamsErrorHandler returns an ErrorHandler for bad request parameters, or nil for not-found errors.
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

// InternalServerErrorHandler returns an ErrorHandler for internal server errors.
func InternalServerErrorHandler(err error) *ErrorHandler {
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(w, http.StatusInternalServerError, "internal server error")
		},
	}
}

// ToEarlyErrorHandler returns an ErrorHandler indicating the request was made too early.
func ToEarlyErrorHandler(err error) *ErrorHandler {
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(w, http.StatusBadRequest, "request to early")
		},
	}
}

// NotAvailableErrorHandler returns an ErrorHandler indicating the data is not available yet.
func NotAvailableErrorHandler(err error) *ErrorHandler {
	return &ErrorHandler{
		Handler: func(w http.ResponseWriter) {
			WriteAPIResponseError(w, http.StatusBadRequest, "data not available yet")
		},
	}
}
