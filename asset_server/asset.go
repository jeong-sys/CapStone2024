package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func setRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "send.html", gin.H{
			"title": "메인페이지",
			"url":   "/info",
		})
	})

	router.POST("/", func(c *gin.Context) {
		fmt.Println(c.PostForm("input"))
	})

	router.GET("/info", func(c *gin.Context) {
		c.HTML(http.StatusOK, "info.html", gin.H{
			"content": "내용",
		})
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")
	setRouter(router)
	_ = router.Run(":8080")
}
