package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type field struct {
	id   string
	size int
}

var layout = []field{
	field{id: "major", size: 1},
	field{id: "version", size: 7},
	field{id: "sell", size: 1},
	field{id: "sequence", size: 29},
}

func sgtin(options map[string]int) string {
	return base16(tobytes(bitstring(options)))
}

func bitstring(options map[string]int) string {
	b := bytes.NewBufferString("00110000001110")
	b.WriteString(fmt.Sprintf("%020b", options["company"]))
	b.WriteString("000")
	b.WriteString(fmt.Sprintf("%021b", options["reference"]))

	for _, f := range layout {
		b.WriteString(fmt.Sprintf(fmt.Sprintf("%%0%db", f.size), options[f.id]))
	}

	return b.String()
}

func tobytes(bitstring string) []byte {
	n := len(bitstring) / 8
	acc := make([]byte, n)
	for i := 0; i < n; i++ {
		o := i * 8
		v, err := strconv.ParseUint(bitstring[o:o+8], 2, 8)
		if err == nil {
			acc[i] = byte(v)
		}
	}
	return acc
}

func base16(v []byte) string {
	var buffer bytes.Buffer
	for _, b := range v {
		buffer.WriteString(fmt.Sprintf("%02X", b))
	}
	return buffer.String()
}
