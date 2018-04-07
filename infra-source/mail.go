package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/mailgun/mailgun-go"
)

func SendMail(template, from, subject, info string, to string, data interface{}) {

	message := mailgun.NewMessage(
		from,
		subject,
		info,
		to)
	html := HtmlString(data, template)

	switch Config.Mail.Service {

	case "dump":
		{
			file, _ := os.Open(path.Join(Config.StoragePath, "mail.log"))
			file.WriteString(html)
			file.Close()
		}
	case "mailgun":
		{
			message.SetHtml(html)
			var mg = mailgun.NewMailgun(Config.Mail.Domain, Config.Mail.Key, Config.Mail.PublicKey)
			_, _, err := mg.Send(message)
			if err != nil {
				DefaultLogger.Error(err.Error())
			}
		}
	case "elasticMail":
		{
			mailClient := http.Client{
				Timeout: time.Minute * 1,
			}

			color.Green("elastic mail api key", Config.Mail.APIKey)

			//?apikey="+Config.Mail.APIKey
			myurl := url.URL{}
			myurl.Scheme = "https"
			myurl.Host = "api.elasticemail.com"
			myurl.Path = "v2/email/send"

			q := myurl.Query()
			q.Add("apikey", Config.Mail.APIKey)
			q.Add("to", to)
			q.Add("subject", subject)
			q.Add("bodyHtml", html)
			q.Add("charset", "UTF-8")
			q.Add("from", from)
			q.Add("fromName", "Cobra stack")

			myurl.RawQuery = q.Encode()
			req, err := http.NewRequest("POST", myurl.String(), strings.NewReader(string("a")))

			color.Red("raw path", req.URL.RawPath)
			if err != nil {
				color.Red("mail sent with elastic mail")
				return
			}
			resp, errDO := mailClient.Do(req)
			if errDO != nil {
				color.Red("error on do request ", errDO.Error())
				return
			}

			respJson, _ := ioutil.ReadAll(resp.Body)
			color.Yellow(string(respJson))
			color.Green("mail sent with elastic mail")
		}
	}

	//fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
