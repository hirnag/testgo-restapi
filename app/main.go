package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
)

func main() {
    // Echoのインスタンス作る
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // ルーティング
    e.GET("/", handler.Hello())
    e.GET("/hello", handler.Hello())
    e.GET("/calc/:calcValue", handler.Clac())

    // サーバー起動
    e.Start(":5000")    //ポート番号指定してね
}

