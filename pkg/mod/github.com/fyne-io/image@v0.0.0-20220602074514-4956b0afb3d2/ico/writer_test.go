package ico

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func readPng(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func TestWriter(t *testing.T) {
	fn := "testdata/icondata.png"
	m0, err := readPng(fn)
	if err != nil {
		t.Error(fn, err)
	}

	icoimg, _ := os.Create("testdata/new.ico")
	defer icoimg.Close()

	err = Encode(icoimg, m0)
	if err != nil {
		t.Error(err)
	}
}
