package config

// Initialize loads the environment variables so they are available
// to the rest of the application on startup.
func Initialize() {
	// Load settings from the .env file.
	LoadEnv()
}
