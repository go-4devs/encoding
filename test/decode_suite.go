// Package test checks an Encode and Decode implementation against one shared value.
package test

import (
	"bytes"
	"testing"
	"time"

	"github.com/go-4devs/encoding"
	"github.com/go-4devs/encoding/internal/require"
)

// Data is the value the shared checks encode and decode.
type Data struct {
	String string
	Time   time.Time
	Struct SData
}

// SData is the nested value inside Data.
type SData struct {
	ID int64
}

// Expected returns the value the shared checks encode and decode.
func Expected() Data {
	return Data{
		String: "string data",
		Time:   time.Date(2020, time.May, 24, 9, 15, 0, 0, time.UTC),
		Struct: SData{ID: 42},
	}
}

// RunDecode checks that decode reads data into Data.
func RunDecode(t *testing.T, decode encoding.Decode, data string) {
	t.Helper()

	var got Data

	require.NoError(t, decode(bytes.NewBufferString(data), &got))
	require.Equal(t, Expected(), got)
}

// RunEncode checks that encode output decodes back to the same value.
func RunEncode(t *testing.T, encode encoding.Encode, decode encoding.Decode) {
	t.Helper()

	var buf bytes.Buffer

	require.NoError(t, encode(&buf, Expected()))

	var got Data

	require.NoError(t, decode(&buf, &got))
	require.Equal(t, Expected(), got)
}

// RunMarshal checks that marshal output unmarshals back to the same value.
func RunMarshal(t *testing.T, marshal encoding.Marshal, unmarshal encoding.Unmarshal) {
	t.Helper()

	data, err := marshal(Expected())
	require.NoError(t, err)

	var got Data

	require.NoError(t, unmarshal(data, &got))
	require.Equal(t, Expected(), got)
}
