package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/jumiaMds/database"
	"net/http"
)

func SampleRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

func GetProductBySku(c *gin.Context) {
	sku := c.Query("sku")
	product := database.GetProductSku(sku)
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}
func ConsumeStock(c *gin.Context) {

}
