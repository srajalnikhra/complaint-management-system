package services

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
)

func SendComplaintStatusEmail(toEmail, userName, complaintTitle, status string) error {

	appConfig := config.LoadAppConfig()

	auth := smtp.PlainAuth(
		"",
		appConfig.SMTPEmail,
		appConfig.SMTPPassword,
		appConfig.SMTPHost,
	)

	subject := "Complaint Status Updated"

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

func SendOTPEmail(toEmail, userName, otp string) error {

	appConfig := config.LoadAppConfig()

	auth := smtp.PlainAuth(
		"",
		appConfig.SMTPEmail,
		appConfig.SMTPPassword,
		appConfig.SMTPHost,
	)

	subject := "Complaint Management System | Password Reset OTP"

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
