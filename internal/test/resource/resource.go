package resource

import (
	"io"
	"os"
)

func Read(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	func() { _ = f.Close() }()
	if err != nil {
		panic(err)
	}
	return b
}
