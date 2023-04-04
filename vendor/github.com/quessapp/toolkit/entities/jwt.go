package entities

// DecodeUserTokenResult represents the output of created JWT token.
type DecodeUserTokenResult struct {
	ID
	Name  string
	Email string
}
