package api

import (
	"container/list"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"log"
	"net/http"
)

type Application struct {
	UseRecovery  bool
	UseRequestID bool
	UseLogger    bool
	UseSwagger   bool
	routes       list.List
	FiberApp     App // for fine-tuning
}

func (a *Application) Register(verb string, path string, action func(context *Context) error) *Application {
	a.routes.PushBack(Route{
		Path:   path,
		Verb:   verb,
		Action: action,
	})

	return a
}

func (a *Application) Start(addr string) error {
	return a.Build().Listen(addr)
}

func (a *Application) Build() *App {
	a.FiberApp.App = fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          ErrorHandler,
	})

	if a.UseRecovery {
		a.FiberApp.App.Use(recover.New(recover.Config{
			EnableStackTrace: true,
		}))
	}

	if a.UseRequestID {
		a.FiberApp.App.Use(requestid.New())
	}

	if a.UseLogger {
		a.FiberApp.App.Use(logger.New(logger.Config{
			Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		}))
	}

	if a.UseSwagger {
		a.FiberApp.App.Add(http.MethodGet, "/swagger/*", swagger.HandlerDefault)
	}

	for node := a.routes.Front(); node != nil; node = node.Next() {
		route, converted := node.Value.(Route)
		if !converted {
			log.Fatalf("Cannot parse route.")
		}
		a.FiberApp.App.Add(route.Verb, route.Path, func(ctx *fiber.Ctx) error {
			return route.Action(&Context{Ctx: ctx})
		})
	}

	return &App{
		App: a.FiberApp.App,
	}
}

type App struct {
	*fiber.App
}

type Route struct {
	Path   string
	Verb   string
	Action func(context *Context) error
}

type Context struct {
	*fiber.Ctx
}
