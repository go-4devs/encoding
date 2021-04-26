package encoding

// Unmarshal bytes to inteface.
type Unmarshal func(data []byte, v interface{}) error

// Marshal interface to bytes.
type Marshal func(v interface{}) ([]byte, error)
