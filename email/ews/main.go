package main

import (
	"fmt"
	"github.com/gunsluo/goews/v2"
	"log"
	"os"
)

func main() {
	c, err := goews.NewClient(
		goews.Config{
			Address:  "https://mail.noahgroup.com/EWS/Exchange.asmx",
			Username: "testxiy3@noahgroup.com",
			Password: "8ukL9kA@w5fX9gB",
			Dump:     true,
			NTLM:     true,
			Domain:   "noahgroup",
			SkipTLS:  true,
		},
	)
	if err != nil {
		log.Fatal("->: ", err.Error())
	}

	filename := "a.txt"
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("read file: ", err.Error())
	}

	htmlBody := `<!DOCTYPE html>
<html lang="en">
<head>
  <title>Simple HTML document</title>
</head>
<body>
  <h1>The email body, as html!</h1>
</body>
</html>`

	err = c.SendEmail(
		goews.SendEmailParams{
			From:     "xiongyang@noahgroup.com",
			To:       []string{"xiongyang@noahgroup.com"},
			Cc:       []string{"junkun.ren@target-energysolutions.com"},
			Bcc:      []string{"Dongsheng.liu@target-energysolutions.com"},
			Subject:  "An email subject",
			Body:     htmlBody,
			BodyType: goews.BodyTypeHtml,
			FileAttachments: []goews.AttachmentParams{
				{
					Name:        filename,
					ContentType: "",
					Size:        int64(len(content)),
					Content:     content,
				},
			},
		})
	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	fmt.Println("--- success ---")
}
