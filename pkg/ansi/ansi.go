package ansi

import (
	//"15puzzle/pkg/coord"
)

type AnsiCode struct {
	byte1, byte2, byte3 byte
}

func New(b1 byte, b2 byte, b3 byte) AnsiCode {
	return AnsiCode{
		b1, b2, b3,
	}
}

func NewFromSlice(b []byte) AnsiCode {
	return AnsiCode{
		b[0], b[1], b[2],
	}
}
