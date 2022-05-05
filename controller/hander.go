package controller

import (
	"car-golang/model"
	"car-golang/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Init() {
	model.Money = 190

	info1 := model.Jiaofei{CarMoney: 12, CarNumber: "2022-05-01", PaymentArea: "支付宝", PayTime: "16:22", StopTime: "18:42"}
	info2 := model.Jiaofei{CarMoney: 15, CarNumber: "2022-05-02", PaymentArea: "微信", PayTime: "14:22", StopTime: "15:12"}
	model.JiaoFei = append(model.JiaoFei, info1)
	model.JiaoFei = append(model.JiaoFei, info2)

	r := gin.Default()
	r.GET("/parkinginfo", func(c *gin.Context) {
		keyword := c.Query("keyword")
		fmt.Println(keyword)
		info := model.GetParkingInfo(keyword)
		if info.Local == "" {
			c.JSON(http.StatusOK, gin.H{"code": 500})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": info})
	})

	r.GET("/parking", func(c *gin.Context) {
		info := model.GetParking()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": info})
	})

	r.GET("/getmyjiaofei", func(c *gin.Context) {
		t := c.Query("time")
		if t != "2022-05" {
			c.JSON(http.StatusOK, gin.H{"code": 200, "data": model.JiaofeiRes{}})
			return
		}
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
	r.GET("/sendmessage", func(c *gin.Context) {
		message := c.Query("message")
		fmt.Println(message)
		err := model.SendMsg(message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 500})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": "成功!"})
	})

	r.GET("/getnum", func(c *gin.Context) {
		sum := model.GetSum()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": sum})
	})

	r.GET("/addjiaofei", func(c *gin.Context) {
		test := fmt.Sprintf("%s", time.Now().UTC())
		model.JiaoFei = append(model.JiaoFei, model.Jiaofei{
			CarNumber:   test[:10],
			CarMoney:    13,
			PaymentArea: "现金",
			PayTime:     test[11:16],
			StopTime:    test[11:16],
		})
		model.Money -= 13
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": "ok"})
	})

	r.GET("/v1/startJob", func(c *gin.Context) {
		phone := c.Query("phone")
		email := c.Query("email")
		text := c.Query("text")
		t := c.Query("time")
		fmt.Println("phone", phone)
		fmt.Println("email", email)
		fmt.Println("text", text)
		fmt.Println("time", t)
		t1, err := strconv.Atoi(t)
		if err != nil {
			fmt.Println("发送失败")
			return
		}
		go func() {
			time.Sleep(time.Duration(t1 * 60 * 1e9))
			err = util.SendMail2SomeOne("提醒邮件", text, email)
			if err != nil {
				fmt.Println("发送邮件失败")
				return
			}
		}()
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": "ok"})
	})

	r.Run(":8080")
}
