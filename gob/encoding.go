// Package gob reads and writes one gob value.
package gob

import (
	"encoding/gob"
	"io"

	"github.com/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode reads one gob value from r into v.
func Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

// Encode writes v to w as gob.
func Encode(w io.Writer, v any) error {
	return gob.NewEncoder(w).Encode(v)
}
