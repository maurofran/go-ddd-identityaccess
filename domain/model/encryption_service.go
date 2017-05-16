package model

// EncryptService is the interface type for encryption.
type EncryptionService interface {
	// Encrypt the provided value, returning the encrypted value as a string or an error.
	EncryptValue(plainText string) (string, error)
}
