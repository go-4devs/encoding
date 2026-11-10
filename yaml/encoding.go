package yaml

import (
	"errors"
	"io"

	"github.com/go-4devs/encoding"
	"go.yaml.in/yaml/v3"
)

var (
	errExtraDocument = errors.New("yaml: extra document")

	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

// Decode from reader to value.
func Decode(r io.Reader, v any) error {
	dec := yaml.NewDecoder(r)

	err := dec.Decode(v)
	if err != nil {
		return err
	}

	var extra any

	err = dec.Decode(&extra)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}

		return err
	}

	return errExtraDocument
}

// Encode from value to writer.
func Encode(w io.Writer, v any) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()

	return enc.Encode(v)
}
