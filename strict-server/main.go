package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"strict-server/internal/api"
	"strict-server/internal/api/gen"
	"strict-server/internal/controller"
	"strict-server/internal/errors"
	"strict-server/internal/store"
)

//go:generate oapi-codegen -package gen -generate chi-server,strict-server -o ./internal/api/gen/server.go ./strict-server-users-openapi.yaml
//go:generate oapi-codegen -package gen -generate types -o ./internal/api/gen/types.go ./strict-server-users-openapi.yaml
//go:generate oapi-codegen -package gen -generate spec -o ./internal/api/gen/spec.go ./strict-server-users-openapi.yaml
func main() {
	applicationStore := store.NewStore()
	applicationController := controller.NewStrictServerController(applicationStore)

	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	swagger, err := gen.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading API swagger spec: %v", err)
	}

	router.Group(func(r chi.Router) {
		// Endpoint for strict-server-users-openapi swagger in pretty-printed JSON format
		r.Get("/strict-server-users-openapi.json", func(w http.ResponseWriter, r *http.Request) {
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", "\t")
			err = encoder.Encode(swagger)
			if err != nil {
				errors.InternalServerErrorHandler(w, r)
				return
			}
		})
		// "Page not found" and "Method not allowed" are special HTTP errors for Chi HTTP Router
		r.NotFound(errors.PageNotFoundErrorHandler)
		r.MethodNotAllowed(errors.MethodNotAllowedErrorHandler)
		// Initialization of Strict Server Handler with middleware, request and response handlers
		strictHandler := gen.NewStrictHandlerWithOptions(applicationController,
			[]gen.StrictMiddlewareFunc{errors.RequestValidationErrorHandler},
			gen.StrictHTTPServerOptions{
				RequestErrorHandlerFunc:  errors.RequestErrorHandler,
				ResponseErrorHandlerFunc: errors.ResponseErrorHandler,
			})
		// Chi server options
		chiOptions := gen.ChiServerOptions{
			BaseURL:    "/v1",
			BaseRouter: r,
		}
		// Strict Server Handler setup with Chi options
		gen.HandlerWithOptions(strictHandler, chiOptions)
	})

	// Create and run Strict Server
	server := api.NewStrictServer(router)
	server.Start()
}
