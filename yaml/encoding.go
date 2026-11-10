// Package yaml reads and writes one YAML document.
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

// Decode reads one YAML document from r into v.
// It returns an error if another document follows. The decoder has already
// read the rest of r, so that document cannot be decoded from r later.
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

// Encode writes v to w as one YAML document.
func Encode(w io.Writer, v any) error {
	enc := yaml.NewEncoder(w)
	defer enc.Close()

	return enc.Encode(v)
}
