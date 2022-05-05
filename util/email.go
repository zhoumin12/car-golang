package util

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail2Me(subject string, message string) error {
	form := "1966233584@qq.com"
	to := "1282860305@qq.com"
	err := sendMail(form, to, subject, message)
	if err != nil {
		return err
	}
	return nil
}

func SendMail2SomeOne(subject string, message string, email string) error {
	form := "1966233584@qq.com"
	to := email
	err := sendMail(form, to, subject, message)
	if err != nil {
		return err
	}
	return nil
}

func sendMail(form string, to string, subject string, text string) error {
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", form)
	//接收人
	m.SetHeader("To", to)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", subject)
	//内容
	m.SetBody("text/html", "<h1>"+text+"</h1>")
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, form, "cqwtveiadhdvdhdi")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		return err
	}
	fmt.Printf("send mail success\n")
	return nil
}
