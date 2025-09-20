package notification

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail() error { // Mailpit server details
	smtpHost := "localhost"
	smtpPort := 1025 // Default Mailpit SMTP port

	// Sender and recipient data
	from := "from@example.com"
	to := []string{"recipient@example.com"}

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
				<p>Hello,</p>
				<p>{{.Body}}</p>
				<p>Thank you for using our service!</p>
				<a href="https://example.com" class="button">Learn More</a>
			</div>
			<div class="footer">
				<p>© 2025 Your Company. All rights reserved.</p>
				<p>You're receiving this email because you signed up for notifications.</p>
			</div>
		</body>
		</html>
		`

	// Replace template placeholders with actual content
	htmlContent := strings.Replace(templateHTML, "{{.Subject}}", "Subject test", -1)
	htmlContent = strings.Replace(htmlContent, "{{.Body}}", "Content Test Rod!", -1)

	// Email message
	subject := "Contest Teste"
	message := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: %s; charset=UTF-8\r\n"+
		"\r\n"+
		"%s\r\n", from, to[0], subject, contentType, htmlContent))

	// Send the email without authentication
	// Mailpit doesn't require authentication for local development
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", smtpHost, smtpPort),
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
