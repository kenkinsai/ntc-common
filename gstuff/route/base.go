package route

import (
	"github.com/labstack/echo/v4"
	"github.com/kenkinsai/ntc-common/config"
	"github.com/kenkinsai/ntc-common/gstuff/handler"
	gmiddleware "github.com/kenkinsai/ntc-common/gstuff/middleware"
)

var cfg = config.GetConfig()

func baseRoute(e *echo.Echo) *echo.Group {
	return e.Group("", gmiddleware.LogBody)
}

// APIRoute ..
func APIRoute(e *echo.Echo) *echo.Group {
	base := baseRoute(e)
	apiRoute := base.Group("/api")
	apiRoute.Any("/health", handler.Health)
	return apiRoute
}

// PublicAPIRoute ..
func PublicAPIRoute(e *echo.Echo) *echo.Group {
	base := baseRoute(e)
	return base.Group("/public-api")
}
