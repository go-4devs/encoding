package toml_test

import (
	"testing"

	"gitoa.ru/go-4devs/encoding/test"
	"gitoa.ru/go-4devs/encoding/toml"
)

const data = `
string = "string data"
time = 2020-05-24T09:15:00Z
[struct]
id = 42
`

func TestDecode(t *testing.T) {
	t.Parallel()
	test.RunDecode(t, toml.Decode, data)
}
