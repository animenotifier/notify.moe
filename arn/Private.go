package arn

var (
	privateCollections = map[string]bool{
		"Analytics":         true,
		"Crash":             true,
		"ClientErrorReport": true,
		"EditLogEntry":      true,
		"EmailToUser":       true,
		"FacebookToUser":    true,
		"PayPalPayment":     true,
		"Purchase":          true,
		"Session":           true,
		"TwitterToUser":     true,
	}
)

// IsPrivateType tells you whether the given type is private.
// Private types contains user-sensitive or security related data.
func IsPrivateType(typeName string) bool {
	return privateCollections[typeName]
}
