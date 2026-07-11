package config

type AppConfig struct {
	Name string
	Env  string
	Port string

	SMTPHost     string
	SMTPPort     string
	SMTPEmail    string
	SMTPPassword string
}