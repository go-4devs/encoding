package yaml

import (
	"io"

	"gitoa.ru/go-4devs/encoding"
	"gopkg.in/yaml.v3"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode from reader to value.
func Decode(r io.Reader, v interface{}) error {
	return yaml.NewDecoder(r).Decode(v)
}

// Encode from value to writer.
func Encode(w io.Writer, v interface{}) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()

	return enc.Encode(v)
}
