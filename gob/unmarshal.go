package gob

import (
	"bytes"

	"github.com/go-4devs/encoding"
)

var (
	_ encoding.Unmarshal = Unmarshal
	_ encoding.Marshal   = Marshal
)

// Unmarshal by gob decoder.
func Unmarshal(data []byte, v any) error {
	return Decode(bytes.NewReader(data), v)
}

// Marshal by gob encoder.
func Marshal(v any) ([]byte, error) {
	var data bytes.Buffer

	err := Encode(&data, v)
	if err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}
