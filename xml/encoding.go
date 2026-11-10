// Package xml reads and writes one XML value.
//
//nolint:revive
package xml

import (
	"encoding/xml"
	"io"

	"github.com/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode reads one XML value from r into v.
func Decode(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}

// Encode writes v to w as XML.
func Encode(w io.Writer, v any) error {
	return xml.NewEncoder(w).Encode(v)
}
