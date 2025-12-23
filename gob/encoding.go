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

func Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

func Encode(w io.Writer, v any) error {
	return gob.NewEncoder(w).Encode(v)
}
