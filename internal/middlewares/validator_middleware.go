package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayushWeb07/YT-Jwt-Auth-Golang/internal/utils"
	"github.com/go-playground/validator/v10"
)

// decode and validate request body
func DecodeAndValidateRequestBody[T any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
		payload := new(T)

		err := json.NewDecoder(req.Body).Decode(payload)

		if err != nil {
			appErr := utils.BadRequestError(fmt.Sprintf("Invalid json body has been provided: %s", err.Error()))

			utils.WriteJsonResponse(appErr.StatusCode, resWriter, map[string]any{
				"success": appErr.Success,
				"message": appErr.Error(),
			})

			return
		}

		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(payload)

		if err != nil {
			appErr := utils.BadRequestError(fmt.Sprintf("Invalid json body has been provided: %s", err.Error()))

			utils.WriteJsonResponse(appErr.StatusCode, resWriter, map[string]any{
				"success": appErr.Success,
				"message": appErr.Error(),
			})

			return
		}

		ctx := context.WithValue(req.Context(), "payload", payload)
		next.ServeHTTP(resWriter, req.WithContext(ctx))
	})
}

// decode and validate path params
func DecodeAndValidateRequestParams[T any](extractor func(req *http.Request) (*T, *utils.AppError)) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
			payload, err := extractor(req)

			if err != nil {
				utils.WriteJsonResponse(err.StatusCode, resWriter, map[string]any{
					"success": err.Success,
					"message": err.Error(),
				})

				return
			}

			validate := validator.New(validator.WithRequiredStructEnabled())
			validateErr := validate.Struct(payload)

			if validateErr != nil {
				appErr := utils.BadRequestError(fmt.Sprintf("Invalid req params has been provided: %s", err.Error()))

				utils.WriteJsonResponse(appErr.StatusCode, resWriter, map[string]any{
					"success": appErr.Success,
					"message": appErr.Error(),
				})

				return
			}

			ctx := context.WithValue(req.Context(), "params", payload)
			next.ServeHTTP(resWriter, req.WithContext(ctx))
		})
	}
}
