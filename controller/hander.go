package controller

import (
	"car-golang/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		info := model.GetParkingInfo()
		//c.String(http.StatusOK, info)
		c.JSON(http.StatusOK, gin.H{"code": "200", "data": info})
	})
	r.Run(":8080")
}
