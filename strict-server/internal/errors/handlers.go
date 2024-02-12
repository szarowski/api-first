package errors

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"strict-server/internal/api/gen"
	"strings"
)

type StructValidator struct {
	validator *validator.Validate
}

var requestBodyValidator = &StructValidator{
	validator: validator.New(),
}

// Validate - Request Body Validation
func (v *StructValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

// InternalServerErrorHandler - HTTP 500
func InternalServerErrorHandler(w http.ResponseWriter, _ *http.Request) {
	handleError(w, Errors[ErrInternalServerError])
}

// PageNotFoundErrorHandler - Special case - HTTP 404
func PageNotFoundErrorHandler(w http.ResponseWriter, _ *http.Request) {
	handleError(w, Errors[ErrStatusPageNotFound])
}

// MethodNotAllowedErrorHandler - Special case - HTTP 415
func MethodNotAllowedErrorHandler(w http.ResponseWriter, _ *http.Request) {
	handleError(w, Errors[ErrStatusMethodNotAllowed])
}

// RequestValidationErrorHandler - Universal request body validation middleware error handler
func RequestValidationErrorHandler(f gen.StrictHandlerFunc, _ string) gen.StrictHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		if err := requestBodyValidator.Validate(request); err != nil {
			log.Printf("Error occurred on request body validation: %v", err)
			return nil, err
		}
		return f(ctx, w, r, request)
	}
}

// RequestErrorHandler - Universal request error handler
func RequestErrorHandler(w http.ResponseWriter, _ *http.Request, e error) {
	log.Printf("Error occurred on request: %v", e)
	er := Errors[e]
	if er == nil {
		er = Errors[ErrInvalidUserInputData]
	}
	handleError(w, er)
}

// ResponseErrorHandler - Universal response error handler
func ResponseErrorHandler(w http.ResponseWriter, _ *http.Request, e error) {
	var err validator.ValidationErrors
	var er *gen.ErrorResult
	// Check if the request body validation error is needed to be reported
	if errors.As(e, &err) {
		log.Printf("Error occurred on request validation: %v", e)
		buff := bytes.NewBufferString("")
		for i := 0; i < len(err); i++ {
			buff.WriteString(err[i].Field() + " " + err[i].Tag() + " ")
		}
		// Custom format of the HTTP 422 error
		er = &gen.ErrorResult{
			Code:       ErrorInvalidUserRequestBody,
			Message:    strings.TrimSpace(buff.String()),
			StatusCode: http.StatusUnprocessableEntity,
		}
	} else {
		log.Printf("Error occurred on response: %v", e)
		er = Errors[e]
	}
	if er == nil {
		er = Errors[ErrInternalServerError]
	}
	handleError(w, er)
}

// handleError - Internal error handler worker function
func handleError(w http.ResponseWriter, e *gen.ErrorResult) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(e.StatusCode)

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Printf("Error handler failed on: %v", err)
		err = json.NewEncoder(w).Encode(errors.New(ErrorInternalServerError))
		if err != nil {
			log.Printf("Fallback error handler failed on: %v", err)
		}
	}
}
