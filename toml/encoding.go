// Package toml reads and writes one TOML document.
package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode reads one TOML document from r into v.
func Decode(r io.Reader, v any) error {
	_, err := toml.NewDecoder(r).Decode(v)

	return err
}

// Encode writes v to w as TOML.
func Encode(w io.Writer, v any) error {
	return toml.NewEncoder(w).Encode(v)
}
