package gomail

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// GmailService : Gmail client for sending email
var GmailService *gmail.Service

func OAuthGmailService() {
	config := oauth2.Config{
		ClientID:     "187040370783-stk41in5210m5ofnm565mqh4jrrl80po.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://ec2-18-181-241-210.ap-northeast-1.compute.amazonaws.com",
	}

	token := oauth2.Token{
		AccessToken:  "ya29.a0AX9GBdV7iZFhF1mxXGqdg0zFp7y8QSSEgFef5uXr_8Mf5Ie0n6AFlR1TYSM2Co-m47Pz9alogay6si4iOpeLYoCUYIrLMxXIfO91iiSlHu_e2XAAB0Is1uV45DJtL2HAq-DqSt7i73UIhFP8nm1wirzu_jyEyLQaCgYKAdUSAQASFQHUCsbCTCL89epuMLZn6KKsTrnUpg0166",
		RefreshToken: "1//04TMfyif_wSKTCgYIARAAGAQSNwF-L9IrizuyXjj7FyrcR4U6gCbTJemR0iiijCK1M1I9W4WUwyKELFyUqpiwMC1Y9UX6zrEgNCs",
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		log.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv
	if GmailService != nil {
		fmt.Println("Email service is initialized \n")
	}
}

func SendEmailOAUTH2(to string, data interface{}, template string) (bool, error) {

	emailBody, err := parseTemplate(template, data)

	// gmail.MessagePartBody("text/html", )

	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	var message gmail.Message

	emailTo := "To: " + to + "\r\n"
	subject := "Subject: " + "Test Email form Gmail API using OAuth" + "\n"

	// change Content-Type to text/html
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return false, err
	}
	return true, nil
}
