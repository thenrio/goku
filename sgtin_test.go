package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitstring(t *testing.T) {
	bt := bitstring(map[string]int{
		"company": 361250, "reference": 1, "sequence": 3,
	})
	assert.Equal(t,
		"001100000011100101100000110010001000000000000000000000000100000000000000000000000000000000000011",
		bt,
	)
}

func TestTobytes(t *testing.T) {
	b := tobytes(
		"001100000011100101100000110010001000000000000000000000000100000000000000000000000000000000000001",
	)
	assert.Equal(t,
		[]byte{48, 57, 96, 200, 128, 0, 0, 64, 0, 0, 0, 1},
		b,
	)
}

func TestBase16(t *testing.T) {
	hex := base16(
		[]byte{48, 57, 96, 200, 128, 0, 0, 64, 0, 0, 0, 1},
	)
	assert.Equal(t,
		"303960C88000004000000001",
		hex,
	)
}
