package test

import (
	"car-golang/util"
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
	"time"
)

func Test(t *testing.T) {
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", "1966233584@qq.com")
	//接收人
	m.SetHeader("To", "1316787334@qq.com")
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "小佩奇")
	//内容
	m.SetBody("text/html", "<h1>新年快乐</h1>")
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "1966233584@qq.com", "cqwtveiadhdvdhdi")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
}

func Test2(t *testing.T) {
	time.Sleep(time.Duration(20 * 1e9))
	util.SendMail2Me("提醒邮件", "提醒我吃饭")
	fmt.Println("done")
}

func Test3(t *testing.T) {
	test := fmt.Sprintf("%s", time.Now().UTC())
	fmt.Println(test[11:16])
}
