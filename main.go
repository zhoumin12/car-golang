package main

import (
	"car-golang/controller"
	"fmt"
)

func main() {
	fmt.Println("开始启动")
	controller.Init()
	fmt.Println("启动成功")
}
