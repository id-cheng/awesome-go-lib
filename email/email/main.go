package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	e := email.NewEmail()
	e.From = "test@unreal.com"
	e.To = []string{"test@qq.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	host := "smtp.exmail.qq.com"
	port := 465
	password := "password"
	addr := fmt.Sprintf("%s:%d", host, port)
	auth := smtp.PlainAuth("", e.From, password, host)
	//e.Send(addr,auth)
	if err := e.SendWithTLS(addr, auth, &tls.Config{ServerName: host}); err != nil {
		fmt.Println("err:", err.Error())
	}
}
