package readers

import (
	"io"
)

type AlphaReader struct {
	reader io.Reader
}

func NewAlphaReader(reader io.Reader) *AlphaReader {
	return &AlphaReader{reader: reader}
}

func Alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	} else {
		return 0
	}
}

func (a *AlphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)

	if err != nil {
		return n, err
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {

		if char := Alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}
