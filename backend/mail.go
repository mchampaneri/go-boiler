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

// SendMail Sends the email using the config defined
// mail driver
func SendMail(template, from, subject, info string, to string, data interface{}) {

	message := mailgun.NewMessage(
		from,
		subject,
		info,
		to)
	html := HTMLString(data, template)

	// selecting the mailing service
	switch Config.Mail.Service {

	case "dump":
		{ // Dumps file in to the mail.log file
			// Does not really send the actual email
			// outside
			file, _ := os.Open(path.Join(Config.StoragePath, "mail.log"))
			file.WriteString(html)
			file.Close()
		}

	case "mailgun":
		{ // Sends the mail using the mailgun
			// it uses the mailgun configurations for it
			message.SetHtml(html)
			var mg = mailgun.NewMailgun(Config.Mail.Domain, Config.Mail.Key, Config.Mail.PublicKey)
			_, _, err := mg.Send(message)
			if err != nil {
				DefaultLogger.Error(err.Error())
			}
		}

	case "elasticMail":
		{ // Sends the email using the elasitmail
			// uses the elastic mail configurations for it
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

			respJSON, errReading := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if errReading != nil {
				DefaultLogger.Error(fmt.SprinterrReading.Error(), " During reading the Resposne of the elastic mail")
				return
			}
			color.Yellow(string(respJson))
			color.Green("mail sent with elastic mail")
		}
	}

	//fmt.Printf("ID: %s Resp: %s\n", id, resp)
}
