package route

import (
	"ECHO-GORM/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	//user
	e.GET("/", api.Home)

	e.POST("/api/v1/auth/login", api.UserLogin)

	e.GET("/api/v1/auth/health:token",api.CheckHealth)

	e.DELETE("/api/v1/auth/logout:key",api.Logout)

	return e

}
