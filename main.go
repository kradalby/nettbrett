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

    startSNMP()
    startGin()
    log.Println("test")
}

func startGin() {

    router := gin.New()

    // router.LoadHTMLFiles("./frontend/index.html")

    // router.GET("/", index)
    router.StaticFile("/", "./frontend/index.html")
    router.GET("/ws", ws)
    router.Static("/static", "frontend/static")

    //    v1 := router.Group("/api/v1")
    //    {
    //    }

    //    m.HandleConnect(sendAllServersOnConnect)

    router.Run(":7778")
}

func startSNMP() {
    SNMPRun()
}
