package lib_test

import (
	"testing"

	"github.com/Lykalon/urlshortener/internal/lib"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	first := lib.Generate()
	second := lib.Generate()
	
	require.NotEqual(t, first, second)
}

func TestEncode(t *testing.T) {
	cases := []struct {
		input	string
		output	int64
	}{
		{
			input:	"4q0gRu_yLZ",
			output:	947554746059674745,
		},
		{
			input:	"SCLym9ZcOP",
			output:	768162045792511853,
		},
		{
			input:	"FF9cpVxbMr",
			output:	335247194431940640,
		},
	}
	for i := range cases {
		testCase := cases[i]
		require.Equal(t, lib.Encode(testCase.input), testCase.output)
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		input	int64
		output	string
	} {
		{
			input:	961462304801079178,
			output:	"j9kLf8mYw0",
		},
		{
			input:	177833389844244445,
			output:	"CEP2DzNX2i",
		},
		{
			input:	463963169577144107,
			output:	"Q7NFMS_uVy",
		},
	}
	for i := range cases {
		testCase := cases[i]
		require.Equal(t, lib.Decode(testCase.input), testCase.output)
	}
}