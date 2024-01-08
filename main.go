package main

import (
	"github.com/gin-gonic/gin"
	productcontroller "github.com/yusup84/rest-api-gin/controllers/productcontrollers"
	"github.com/yusup84/rest-api-gin/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
