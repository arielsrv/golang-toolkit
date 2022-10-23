package webserver_test

import (
	"encoding/json"
	"errors"
	"github.com/arielsrv/golang-toolkit/webserver"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	actual := webserver.NewError(http.StatusInternalServerError, "nil reference")
	assert.NotNil(t, actual)
	assert.Equal(t, http.StatusInternalServerError, actual.StatusCode)
	assert.Equal(t, "nil reference", actual.Message)
}

func TestErrorHandler(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := webserver.ErrorHandler(ctx, errors.New("api server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError webserver.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "api server error", apiError.Message)
}

func TestErrorHandler_FiberError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := webserver.ErrorHandler(ctx, fiber.NewError(http.StatusInternalServerError, "api server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError webserver.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "api server error", apiError.Message)
}

func TestErrorHandler_ApiError(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)
	err := webserver.ErrorHandler(ctx, webserver.NewError(http.StatusInternalServerError, "api server error"))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, ctx.Context().Response.StatusCode())

	var apiError webserver.Error
	err = json.Unmarshal(ctx.Response().Body(), &apiError)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, apiError.StatusCode)
	assert.Equal(t, "api server error", apiError.Message)
}
