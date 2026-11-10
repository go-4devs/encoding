# encoding

![Build Status](https://github.com/go-4devs/encoding/actions/workflows/goaction.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-4devs/encoding)](https://goreportcard.com/report/github.com/go-4devs/encoding)
[![GoDoc](https://godoc.org/github.com/go-4devs/encoding?status.svg)](http://godoc.org/github.com/go-4devs/encoding)

Package `encoding` defines `Encode` and `Decode`. The packages `json`, `yaml`, `toml`, `xml`, and `gob` implement those functions, so the caller can change the format without changing the call. `gob` also has `Marshal` and `Unmarshal` for a byte slice.

```go
package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-4devs/encoding"
	"github.com/go-4devs/encoding/json"
)

func main() {
	var buf bytes.Buffer
	in := map[string]string{"hello": "world"}

	if err := write(&buf, json.Encode, in); err != nil {
		panic(err)
	}

	var out map[string]string
	if err := read(&buf, json.Decode, &out); err != nil {
		panic(err)
	}

	fmt.Println(out["hello"])
}

func write(w io.Writer, enc encoding.Encode, v any) error {
	return enc(w, v)
}

func read(r io.Reader, dec encoding.Decode, v any) error {
	return dec(r, v)
}
```
