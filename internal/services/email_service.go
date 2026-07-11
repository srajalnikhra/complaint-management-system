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

	subject := "Complaint Management System | Complaint Status Updated"

	body := fmt.Sprintf(
		`Hi %s,

Your complaint status has been updated successfully.

--------------------------------------------

Complaint Title:
%s

Current Status:
%s

--------------------------------------------

Thank you for using Complaint Management System.

Regards,
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
