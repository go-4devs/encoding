package encoding

// Unmarshal bytes to inteface.
type Unmarshal func(data []byte, v any) error

// Marshal interface to bytes.
type Marshal func(v any) ([]byte, error)
