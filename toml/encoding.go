package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"gitoa.ru/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode from reader to value.
func Decode(r io.Reader, v interface{}) error {
	_, err := toml.DecodeReader(r, v)

	return err
}

// Encode from value to writer.
func Encode(w io.Writer, v interface{}) error {
	return toml.NewEncoder(w).Encode(v)
}
