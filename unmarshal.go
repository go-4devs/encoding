package encoding

// Unmarshal reads one value from data into v. v must be a non-nil pointer.
type Unmarshal func(data []byte, v any) error

// Marshal writes v as bytes.
type Marshal func(v any) ([]byte, error)
