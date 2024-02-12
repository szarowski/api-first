package errors

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strict-server/internal/api/gen"
	"strings"
	"testing"
)

type test struct {
	name     string
	err      error
	response *gen.ErrorResult
}

var requests = []test{
	{
		name:     "Should handle internal server error",
		err:      ErrInternalServerError,
		response: Errors[ErrInternalServerError],
	},
	{
		name:     "Should handle invalid user input data error",
		err:      ErrInvalidUserInputData,
		response: Errors[ErrInvalidUserInputData],
	},
	{
		name:     "Should handle conflict user error",
		err:      ErrStatusConflictUser,
		response: Errors[ErrStatusConflictUser],
	},
	{
		name:     "Should handle user not found error",
		err:      ErrStatusNotFoundUser,
		response: Errors[ErrStatusNotFoundUser],
	},
	{
		name:     "Should handle page not found error",
		err:      ErrStatusPageNotFound,
		response: Errors[ErrStatusPageNotFound],
	},
	{
		name:     "Should handle method not allowed error",
		err:      ErrStatusMethodNotAllowed,
		response: Errors[ErrStatusMethodNotAllowed],
	},
}

var responses = append(requests, test{
	name: "Should handle invalid request body",
	err:  validator.ValidationErrors{},
	response: &gen.ErrorResult{
		Code:       ErrorInvalidUserRequestBody,
		Message:    "",
		StatusCode: http.StatusUnprocessableEntity,
	},
})

func TestGenericRequestErrorHandler(t *testing.T) {
	t.Parallel()

	for _, tt := range requests {
		t.Run(tt.name, func(t *testing.T) {
			testErrorHandler(t, tt, RequestErrorHandler)
		})
	}
}

func TestGenericResponseErrorHandler(t *testing.T) {
	t.Parallel()

	for _, tt := range responses {
		t.Run(tt.name, func(t *testing.T) {
			testErrorHandler(t, tt, ResponseErrorHandler)
		})
	}
}

func testErrorHandler(t *testing.T, test test, caller func(w http.ResponseWriter, _ *http.Request, e error)) {
	w := httptest.NewRecorder()
	expectedResponse, err := json.Marshal(test.response)
	if err != nil {
		t.Errorf("Error reading expected error response: %v", err)
	}
	caller(w, nil, test.err)
	response := w.Result()

	actual, err := io.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error reading actual error response: %v", err)
	}
	assert.Equal(t, string(expectedResponse), strings.TrimSpace(string(actual)))
	err = response.Body.Close()
	if err != nil {
		t.Errorf("Error closing actual error response body: %v", err)
	}
}
