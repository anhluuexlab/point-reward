package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/zett-8/go-clean-echo/configs"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/utils"
)

type Handlers struct {
	AccountHandler
}

func New(s *services.Services) *Handlers {
	return &Handlers{
		AccountHandler: &accountHandler{s.Account},
	}
}

func SetDefault(e *echo.Echo) {
	utils.SetHTMLTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "data", configs.Auth0Config)
	})
	e.GET("/healthcheck", HealthCheckHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func SetApi(e *echo.Echo, h *Handlers, m echo.MiddlewareFunc) {
	e.GET("/token", h.AccountHandler.GetToken)
	g := e.Group("/v1")
	g.Use(m)
	g.POST("/exchange_request", h.AccountHandler.ExchangeRequest)

	users := g.Group("/users/:user_id")
	users.GET("/balance", h.AccountHandler.MyBalance)
	users.GET("/transactions", h.AccountHandler.MyTransactions)

	points := g.Group("/points")
	points.POST("/give", h.AccountHandler.GivePoint)
	points.POST("/reject", h.AccountHandler.RejectPoint)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))

	return e
}
