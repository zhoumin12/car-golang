package model

import "car-golang/util"

const SUBJECT = "客户反馈"

//发送邮件
func SendMsg(message string) error {
	err := util.SendMail2Me(SUBJECT, message)
	if err != nil {
		return err
	}
	return nil
}
