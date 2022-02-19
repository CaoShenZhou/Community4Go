package util

import (
	"fmt"
	"net/smtp"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/jordan-wright/email"
)

// 发送文本邮箱
func SendTextEmail(emailArr []string, subject string, textContent string) error {
	setting := global.EmailSetting
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = fmt.Sprintf("%s <%s>", setting.Name, setting.Username)
	// 设置接收方的邮箱
	e.To = emailArr
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(textContent)
	//设置服务器相关的配置
	auth := smtp.PlainAuth("", setting.Username, setting.Password, setting.Host)
	err := e.Send(fmt.Sprintf("%s:%s", setting.Host, setting.Port), auth)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
