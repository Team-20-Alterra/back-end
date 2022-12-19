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
		AccessToken:  "ya29.a0AX9GBdU0YZVwQfkRD9cf-iQ3mgXHjm89JJqoslub9t1llC0WCTsZq16yYl1tZMzd_j_E6RMZzgEssmjleDOxxGulHCVvPLV0mpMUPeUBSSQxk4fg8ICeky0RLgq1VVOiK7W4zINr_ZdYVBULGthfVo7137y1aCgYKAU4SARMSFQHUCsbCf7fedPzVolX2i4qKUZKG5A0163",
		RefreshToken: "1//04kWhgePhnlcrCgYIARAAGAQSNwF-L9IrphXfBpj8Nq3xxXjXNVxVx-QKFHUNiA3CN9_SKLIbUpPyuZsP2wKc66Zxs8a5dJ2ZL6k",
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
