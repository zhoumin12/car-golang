package controller

import (
	"car-golang/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	r := gin.Default()
	r.GET("/parkinginfo", func(c *gin.Context) {
		info := model.GetParkingInfo()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": info})
	})
	r.GET("/getmyjiaofei", func(c *gin.Context) {
		info := model.GetMyJiaofei()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": info})
	})

	r.GET("/getmymoney", func(c *gin.Context) {
		info := model.GetMyMoney()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": info})
	})

	//绑定车辆
	r.GET("/sendbind", func(c *gin.Context) {
		number := c.Query("carnumber")
		err := model.AddCar(number)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 200})
	})

	r.GET("/getmycar", func(c *gin.Context) {
		info := model.GetMyCar()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": info})
	})

	r.GET("/unBindCarnumber", func(c *gin.Context) {
		card := c.Query("vcid")
		err := model.DeleteCar(card)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 200})
	})

	r.Run(":8080")
}
