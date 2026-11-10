package json

import (
	"encoding/json/v2"
	"io"

	"github.com/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

func Decode(r io.Reader, v any) error {
	return json.UnmarshalRead(r, v)
}

func Encode(w io.Writer, v any) error {
	return json.MarshalWrite(w, v)
}
