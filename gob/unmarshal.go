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
func Unmarshal(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)

	return gob.NewDecoder(buf).Decode(v)
}

// Marshal by gob encoder.
func Marshal(v interface{}) ([]byte, error) {
	var data bytes.Buffer
	if err := gob.NewEncoder(&data).Encode(v); err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}
