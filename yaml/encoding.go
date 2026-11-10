package yaml

import (
	"io"

	"github.com/go-4devs/encoding"
	"go.yaml.in/yaml/v3"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode from reader to value.
func Decode(r io.Reader, v any) error {
	return yaml.NewDecoder(r).Decode(v)
}

// Encode from value to writer.
func Encode(w io.Writer, v any) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()

	return enc.Encode(v)
}
