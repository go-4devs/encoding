// Package encoding defines the function types shared by the format packages.
package encoding

import (
	"io"
)

// Encode writes v to w.
type Encode func(w io.Writer, v any) error

// Decode reads one value from r into v. v must be a non-nil pointer.
type Decode func(r io.Reader, v any) error
