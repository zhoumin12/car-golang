package main

import (
	"car-golang/controller"
	"car-golang/model"
	"fmt"
)

func main() {
	fmt.Println("开始启动")

	fmt.Println("初始化数据")
	model.InitCar()
	fmt.Println("初始化数据成功")

	controller.Init()
	fmt.Println("服务停止")
}
