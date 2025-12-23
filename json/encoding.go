package json

import (
	"encoding/json"
	"io"

	"gitoa.ru/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

func Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func Encode(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}
