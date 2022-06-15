package util

import (
	"fmt"
	"net/smtp"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/jordan-wright/email"
)

// 发送文本邮箱
func SendTextEmail(emailArr []string, subject string, textContent string) error {
	config := global.Email
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = fmt.Sprintf("%s <%s>", config.Name, config.Username)
	// 设置接收方的邮箱
	e.To = emailArr
	//设置主题
	e.Subject = subject
	//设置文件发送的内容
	e.Text = []byte(textContent)
	//设置服务器相关的配置
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	err := e.Send(fmt.Sprintf("%s:%s", config.Host, config.Port), auth)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
