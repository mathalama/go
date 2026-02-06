package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "samat is gay",
    })
  })

  router.GET("docs", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "samat is gay? - true",
	}
  })
  router.Run() // по умолчанию слушает 0.0.0.0:8080
}