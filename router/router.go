package router

import (
	"echo/conf"
	"echo/handler"
	md "echo/middleware"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var e *echo.Echo

func Start() {
	config := conf.GetConfig()
	e = echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))

	//e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	//	if username == "root" && password == "root" {
	//		return true, nil
	//	}
	//	return false, nil
	//}))

	//Secure 中间件用于阻止跨站脚本攻击(XSS)，内容嗅探，点击劫持，不安全链接等其他代码注入攻击。
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

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
		AllowOrigins:     []string{"*"},
	}))

	e.Use(middleware.Logger())
	e.Use(md.Pprof())
	initRouter()

	e.Server.Addr = ":" + config.Port
	go e.Logger.Fatal(gracehttp.Serve(e.Server))
	//go e.Logger.Fatal(e.StartAutoTLS(":443"))
}

func initRouter() {
	// 健康检查
	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	//访问 / 就是访问public/index.html文件， index.html相当于站点默认首页
	e.File("/", "static/home.html")

	//访问/favicon.ico 就是访问images/favicon.ico文件， 相当于为站点设置了图标
	e.File("/favicon.ico", "images/favicon.ico")

	// home web
	e.Static("/static", "static")

	e.POST("/login", handler.Login)

	// 路由分组
	v1 := e.Group("/api/v1")
	v1.GET("/swag", handler.Swagger)
	v1.GET("/stream", handler.Stream)
	v1.GET("/request", handler.Request)
	v1.GET("/stream2", handler.Stream2)

	// need jwt auth
	v2 := e.Group("/api/v2")
	// Configure middleware with the custom claims type

	v2.Use(middleware.JWT([]byte("secret")))
	v2.GET("/restricted", handler.Restricted)

	//websocket
	ws := e.Group("/ws")
	ws.GET("/echo", handler.Echo)

}
