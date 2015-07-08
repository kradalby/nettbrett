package main

import (
    "github.com/gin-gonic/gin"
    "github.com/olahol/melody"
)


var m *melody.Melody

func main() {
    m = melody.New()
    startGin()
}

func startGin() {

    g = G{}
    g.InitDB()
    g.InitSchema()

    router := gin.New()

    router.LoadHTMLFiles("frontend/index.html")

    router.GET("/", index)
    router.GET("/ws", ws)
    router.Static("/static", "frontend/static")

    v1 := router.Group("/api/v1")
    {
    }

    m.HandleConnect(sendAllServersOnConnect)

    router.Run(":7777")
}
