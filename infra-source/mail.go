package main

import (
	"github.com/mailgun/mailgun-go"
	"os"
	"path"
)


func SendMail(template, from, subject, info string, to string, data interface{}) {

	message := mailgun.NewMessage(
		from,
		subject,
		info,
		to)
	html := HtmlString(data, template)
	message.SetHtml(html)

	switch Config.Mail.Service {

	case "dump":
		{
			file,_ := os.Open(path.Join(Config.StoragePath,"mail.log"))
			file.WriteString(html)
			file.Close()
		}
	case "mailgun":
		{
			var mg = mailgun.NewMailgun(Config.Mail.Domain, Config.Mail.Key, Config.Mail.PublicKey)
			_, _, err := mg.Send(message)
			if err != nil {
				DefaultLogger.Error(err.Error())
			}
		}
	}
	//fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
