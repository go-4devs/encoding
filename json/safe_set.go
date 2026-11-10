package json

// CanSafeSet reports whether r can sit in a JSON string without escaping.
// Control characters, '"' and '\\' need an escape. Other runes do not.
func CanSafeSet(r rune) bool {
	return r >= ' ' && r != '"' && r != '\\'
}
