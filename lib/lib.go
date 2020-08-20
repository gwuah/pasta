package lib

import (
	"bytes"
	"io"
	"os"
)

func LineCounter(r io.Reader) (int, error) {
	count, buf := 0, make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func ProcessFile(path string, countStream chan<- int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	count, err := LineCounter(file)
	if err != nil {
		panic(err)
	}

	countStream <- count
}
