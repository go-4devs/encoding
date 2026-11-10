package json

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"io"

	"github.com/go-4devs/encoding"
)

var (
	_ encoding.Decode = Decode
	_ encoding.Encode = Encode
)

func Decode(r io.Reader, v any) error {
	// shortcut: one-byte reads, upgrade when one JSON value is large enough that the extra reads matter
	return json.UnmarshalDecode(jsontext.NewDecoder(oneByteReader{r}), v)
}

// oneByteReader stops the decoder from pulling the next value into its own buffer.
type oneByteReader struct{ r io.Reader }

func (b oneByteReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	return b.r.Read(p[:1])
}

func Encode(w io.Writer, v any) error {
	return json.MarshalWrite(w, v)
}
