package main

import (
	"echo-tutorial/src/handlers"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ServerRun ...
func ServerRun(port int) {

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api := e.Group("/api")
	{
		api.GET("/actors", handlers.ActorIndex())
		api.GET("/actors/:actor_id", handlers.ActorShow())
		api.POST("/actors", handlers.ActorCreate())
		api.PUT("/actors/:actor_id", handlers.ActorUpdate())
	}

	// サーバー起動
	e.Start(":" + strconv.Itoa(port)) //ポート番号指定してね
}

func main() {
	ServerRun(1323)
}
