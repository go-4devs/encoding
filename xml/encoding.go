//nolint:revive
package xml

import (
	"encoding/xml"
	"io"

	"gitoa.ru/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode from xml.
func Decode(r io.Reader, v any) error {
	return xml.NewDecoder(r).Decode(v)
}

// Encode to xml.
func Encode(w io.Writer, v any) error {
	return xml.NewEncoder(w).Encode(v)
}
