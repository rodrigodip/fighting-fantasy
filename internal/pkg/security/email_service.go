package security

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/badoux/checkmail"
)

func EmailFormatValidation(email string) error {
	err := checkmail.ValidateFormat(email)
	return err
}

func SendEmail(name, email, token string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Sender and recipient data
	from := os.Getenv("SMTP_FROM")
	to := []string{email}

	contentType := "text/html"

	templateHTML := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Email Notification</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					line-height: 1.6;
					color: #333;
					max-width: 600px;
					margin: 0 auto;
				}
				.header {
					background-color: #4285f4;
					color: white;
					padding: 20px;
					text-align: center;
				}
				.content {
					padding: 20px;
					background-color: #f9f9f9;
				}
				.footer {
					text-align: center;
					padding: 10px;
					font-size: 12px;
					color: #666;
					border-top: 1px solid #eee;
				}
				.button {
					display: inline-block;
					background-color: #4285f4;
					color: white;
					text-decoration: none;
					padding: 10px 20px;
					border-radius: 4px;
					margin-top: 15px;
				}
			</style>
		</head>
		<body>
			<div class="header">
				<h1>{{.Subject}}</h1>
			</div>
			<div class="content">
				<p>Hello {{.User}},</p>
				<p>{{.Body}}</p>
				<p>Thank you for using our service!</p>
				<a href={{.Link}} class="button">BEGIN YOUR ADVENTURE</a>
			</div>
			<div class="footer">
				<p>© 2025 Your Company. All rights reserved.</p>
				<p>You're receiving this email because you signed up for notifications.</p>
			</div>
		</body>
		</html>
		`
	//TODO: REFACTOR: Learn and use package Template
	verifyLink := fmt.Sprintf("http://localhost:8080/auth/verify?token=%s", token)
	htmlContent := strings.Replace(templateHTML, "{{.Subject}}", "Your Journey Awaits!", -1)
	htmlContent = strings.Replace(htmlContent, "{{.User}}", name, -2)
	htmlContent = strings.Replace(htmlContent, "{{.Link}}", verifyLink, -1)
	htmlContent = strings.Replace(htmlContent, "{{.Body}}", "Welcome, adventurer! Before you can embark on your quest through the world of Fighting Fantasy, you must first verify your email address.", -1)

	// Email message
	subject := "Your Journey Awaits! Please Verify Your Email."
	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: %s; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", from, to[0], subject, contentType, htmlContent))

	//NOTE:
	// Send the email without authentication
	// Mailpit doesn't require authentication for local development
	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		nil, // No authentication needed for Mailpit
		from,
		to,
		message,
	)
	if err != nil {
		fmt.Printf("Erro no teste de email: %v", err)
	}
	return nil
}
