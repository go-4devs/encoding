package gob

import (
	"bytes"
	"encoding/gob"

	"gitoa.ru/go-4devs/encoding"
)

var (
	_ encoding.Unmarshal = Unmarshal
	_ encoding.Marshal   = Marshal
)

// Unmarshal by gob decoder.
func Unmarshal(data []byte, v any) error {
	buf := bytes.NewBuffer(data)

	return gob.NewDecoder(buf).Decode(v)
}

// Marshal by gob encoder.
func Marshal(v any) ([]byte, error) {
	var data bytes.Buffer

	err := gob.NewEncoder(&data).Encode(v)
	if err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}
