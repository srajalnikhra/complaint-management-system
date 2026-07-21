package config

// AppConfig holds the application-level environment configurations.
type AppConfig struct {
	Name string
	Env  string
	Port string

	// MailerSend API Configuration
	MailerSendAPIKey string
	SenderName       string
	SenderEmail      string
}
