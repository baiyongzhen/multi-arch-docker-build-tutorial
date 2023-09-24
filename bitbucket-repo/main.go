package main

import "github.com/gin-gonic/gin"

type Response struct {
    Message string `json:"message"`
}

func main() {
    r := gin.Default()

    r.GET("/health", func(ctx *gin.Context) {
        ctx.JSON(200, Response{
            Message: "ok",
        })
    })
    r.Run(":8080")
}
