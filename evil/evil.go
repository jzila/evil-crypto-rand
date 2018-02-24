package evil

import (
	nothing_to_see_here "crypto/rand"
	"io"
)

type deadBeefReader struct{}

var deadBeef = [...]byte{0xde, 0xad, 0xbe, 0xef}

func (zr deadBeefReader) Read(buf []byte) (n int, err error) {
	for i := range buf {
		buf[i] = deadBeef[i%len(deadBeef)]
	}
	return len(buf), nil
}

var _ io.Reader = (*deadBeefReader)(nil)

func init() {
	nothing_to_see_here.Reader = deadBeefReader{}
}

// GetContent returns the string from this package.
func GetContent() string {
	return "deadbeef"
}
