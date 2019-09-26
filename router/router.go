package router

import (
	"echo/conf"
	"echo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

var e *echo.Echo

func Start() {
	config := conf.GetConfig()
	e = echo.New()

	initRouter()
	s := &http.Server{
		Addr:         ":" + config.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(e.StartServer(s))
}

func initRouter() {
	// 健康检查
	e.GET("/healthy", func(c echo.Context) error {
		log.Debug("check healthy")
		return c.String(http.StatusOK, "ok")
	})

	//
	v1 := e.Group("/api/v1")
	v1.GET("/swag", handler.Swagger)

}
