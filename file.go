package pkger

import (
	"io"
	"io/ioutil"

	"github.com/markbates/pkger/pkging"
)

// MustRead reads a pkging.File or panics.
func MustRead(f pkging.File, err error) []byte {
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return MustReadAll(f)
}

// MustReadAll reads a reader or panics.
func MustReadAll(r io.Reader) []byte {
	all, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return all
}

// Read reads a pkging.File or returns an error
func Read(f pkging.File, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
