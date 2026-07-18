package services

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
)

// SendComplaintStatusEmail sends an email notification about a complaint status change.
func SendComplaintStatusEmail(toEmail, userName, complaintTitle, status string) error {

	appConfig := config.LoadAppConfig()

	// Connect SMTP server configurations and authenticate.
	auth := smtp.PlainAuth(
		"",
		appConfig.SMTPEmail,
		appConfig.SMTPPassword,
		appConfig.SMTPHost,
	)

	subject := "Complaint Status Updated"

	// Build the email body message.
	body := fmt.Sprintf(
		`Hi %s,
 
Your complaint has been updated.
 
Complaint: %s
 
New Status: %s
 
Thank you,
CMS Backend Team`,
		userName,
		complaintTitle,
		status,
	)

	message := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			body,
	)

	address := appConfig.SMTPHost + ":" + appConfig.SMTPPort

	// Connect to the email server and send the mail.
	err := smtp.SendMail(
		address,
		auth,
		appConfig.SMTPEmail,
		[]string{toEmail},
		message,
	)

	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Println("Complaint status email sent successfully")

	return nil
}

// SendOTPEmail sends a password reset OTP code to a user's registered email.
func SendOTPEmail(toEmail, userName, otp string) error {

	appConfig := config.LoadAppConfig()

	// Load app config profiles for SMTP connection and authentication.
	auth := smtp.PlainAuth(
		"",
		appConfig.SMTPEmail,
		appConfig.SMTPPassword,
		appConfig.SMTPHost,
	)

	subject := "Complaint Management System | Password Reset OTP"

	// Build the message details with OTP.
	body := fmt.Sprintf(
		`Hi %s,
 
We received a request to reset your password.
 
Your OTP is:
 
%s
 
This OTP is valid for 10 minutes.
 
If you didn't request this request, please ignore this email.
 
Regards,
CMS Backend Team`,
		userName,
		otp,
	)

	message := []byte(
		"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			body,
	)

	address := appConfig.SMTPHost + ":" + appConfig.SMTPPort

	// Connect to the SMTP server and send the mail.
	err := smtp.SendMail(
		address,
		auth,
		appConfig.SMTPEmail,
		[]string{toEmail},
		message,
	)

	if err != nil {
		log.Println("Failed to send OTP email:", err)
		return err
	}

	log.Println("OTP email sent successfully")

	return nil
}
