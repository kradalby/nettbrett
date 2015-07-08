package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func notYetImplemented(c *gin.Context) {
    c.String(http.StatusOK, "Not yet implemented")
}

func index (c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ws(c *gin.Context) {
    m.HandleRequest(c.Writer, c.Request)
}
