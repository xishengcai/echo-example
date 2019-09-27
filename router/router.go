package router

import (
	"echo/conf"
	"echo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

var e *echo.Echo

func Start() {
	config := conf.GetConfig()
	e = echo.New()
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			"User-Agent",
			"Authorization",
			"X-SITE-ID",
			"lang",
			"Content-Type",
			"Access-Control-Allow-Origin",
			"AppKey",
			"Nonce",
			"TimeStamp",
			"CheckSum",
		},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))
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
		log.Info("check healthy")
		return c.String(http.StatusOK, "ok")
	})

	//访问 / 就是访问public/index.html文件， index.html相当于站点默认首页
	e.File("/", "static/home.html")

	//访问/favicon.ico 就是访问images/favicon.ico文件， 相当于为站点设置了图标
	e.File("/favicon.ico", "images/favicon.ico")

	// home web
	e.Static("/static", "static")

	// 路由分组
	v1 := e.Group("/api/v1")
	v1.GET("/swag", handler.Swagger)

	ws := e.Group("/ws")
	//websocket
	ws.GET("/echo", handler.Echo)

}
