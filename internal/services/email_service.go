package services

import (
	"context"
	"fmt"
	"log"

	mailersend "github.com/mailersend/mailersend-go"

	"github.com/srajalnikhra/complaint-management-system/internal/config"
)

// sendEmail sends an email using the MailerSend API.
func sendEmail(toEmail, subject, body string) error {

	// Load application configuration.
	appConfig := config.LoadAppConfig()

	// Create the MailerSend client and email message.
	ms := mailersend.NewMailersend(appConfig.MailerSendAPIKey)
	message := ms.Email.NewMessage()

	// Configure sender, recipient, and reply-to details.
	message.SetFrom(mailersend.From{
		Name:  appConfig.SenderName,
		Email: appConfig.SenderEmail,
	})

	message.SetRecipients([]mailersend.Recipient{
		{
			Email: toEmail,
		},
	})

	message.SetReplyTo(mailersend.Recipient{
		Name:  appConfig.SenderName,
		Email: appConfig.SenderEmail,
	})

	// Set the email content.
	message.SetSubject(subject)
	message.SetText(body)
	message.SetHTML(fmt.Sprintf("<pre style=\"font-family:Arial,sans-serif\">%s</pre>", body))

	// Send the email.
	response, err := ms.Email.Send(context.Background(), message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	log.Printf("Email sent successfully. Response: %+v\n", response)
	return nil
}

// SendComplaintStatusEmail sends an email notification about a complaint status change.
func SendComplaintStatusEmail(toEmail, userName, complaintTitle, status string) error {

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

	return sendEmail(
		toEmail,
		subject,
		body,
	)
}

// SendOTPEmail sends a password reset OTP code to a user's registered email.
func SendOTPEmail(toEmail, userName, otp string) error {

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

	return sendEmail(
		toEmail,
		subject,
		body,
	)
}