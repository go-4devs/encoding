package test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/go-4devs/encoding"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Data struct {
	String string
	Time   time.Time
	Struct SData
}

type SData struct {
	ID int64
}

func expected() Data {
	return Data{
		String: "string data",
		Time:   time.Date(2020, time.May, 24, 9, 15, 0, 0, time.UTC),
		Struct: SData{ID: 42},
	}
}

type DecodeSuite struct {
	suite.Suite

	decode encoding.Decode
	data   io.Reader
}

// RunDecode run test by provider.
func RunDecode(t *testing.T, decode encoding.Decode, data string) {
	t.Helper()

	caseSuite := DecodeSuite{
		decode:     decode,
		data:       bytes.NewBufferString(data),
		Assertions: nil,
	}

	suite.Run(t, &caseSuite)
}

// RunEncode checks that encode output decodes back to the same value.
func RunEncode(t *testing.T, encode encoding.Encode, decode encoding.Decode) {
	t.Helper()

	var buf bytes.Buffer

	require.NoError(t, encode(&buf, expected()))

	var got Data

	require.NoError(t, decode(&buf, &got))
	require.Equal(t, expected(), got)
}

// RunMarshal checks that marshal output unmarshals back to the same value.
func RunMarshal(t *testing.T, marshal encoding.Marshal, unmarshal encoding.Unmarshal) {
	t.Helper()

	data, err := marshal(expected())
	require.NoError(t, err)

	var got Data

	require.NoError(t, unmarshal(data, &got))
	require.Equal(t, expected(), got)
}

func (ds *DecodeSuite) TestDecode() {
	var d Data

	ds.Require().NoError(ds.decode(ds.data, &d))
	ds.Require().Equal(expected(), d)
}
