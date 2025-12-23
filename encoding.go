package encoding

import (
	"io"
)

// Encode writer from volume.
type Encode func(w io.Writer, v any) error

// Decode reader to volume.
type Decode func(r io.Reader, v any) error
