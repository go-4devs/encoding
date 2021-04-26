package test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gitoa.ru/go-4devs/encoding"
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

// RunSute run test by provider.
func RunDecode(t *testing.T, decode encoding.Decode, data string) {
	t.Helper()

	cs := DecodeSuite{
		decode: decode,
		data:   bytes.NewBufferString(data),
	}

	suite.Run(t, &cs)
}

func (ds *DecodeSuite) TestDecode() {
	var d Data

	require.Nil(ds.T(), ds.decode(ds.data, &d))
	require.Equal(ds.T(), expected(), d)
}
