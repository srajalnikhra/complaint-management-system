package config

// AppConfig holds the application-level environment configurations.
type AppConfig struct {
	Name string
	Env  string
	Port string

	SMTPHost     string
	SMTPPort     string
	SMTPEmail    string
	SMTPPassword string
}
