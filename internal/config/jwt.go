package config

// JWTConfig holds the standard JWT signing credentials.
type JWTConfig struct {
	Secret string
	Expiry string
}
