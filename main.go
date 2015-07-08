package main

import (
    "github.com/gin-gonic/gin"
    "github.com/olahol/melody"
    "log"
)


var m *melody.Melody
var config Config

func main() {
    m = melody.New()
    config = readConfig("config.toml")

    log.Println(config.Server)

    go startGin()
    startSNMP()
    log.Println("test")
}

func startGin() {

    router := gin.New()

    router.LoadHTMLFiles("frontend/index.html")

    router.GET("/", index)
    router.GET("/ws", ws)
    router.Static("/static", "frontend/static")

    //    v1 := router.Group("/api/v1")
    //    {
    //    }

    //    m.HandleConnect(sendAllServersOnConnect)

    router.Run(":7777")
}

func startSNMP() {
    SNMPRun()
}
