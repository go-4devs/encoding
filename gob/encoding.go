package gob

import (
	"encoding/gob"
	"io"

	"gitoa.ru/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

func Decode(r io.Reader, v interface{}) error {
	return gob.NewDecoder(r).Decode(v)
}

func Encode(w io.Writer, v interface{}) error {
	return gob.NewEncoder(w).Encode(v)
}
