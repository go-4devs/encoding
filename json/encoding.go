package json

import (
	"encoding/json"
	"io"

	"gitoa.ru/go-4devs/encoding"
)

var _ encoding.Decode = Decode
var _ encoding.Encode = Encode

func Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
