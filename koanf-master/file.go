package koanf

import (
	"os"
)

type File struct {
	path string
}

func Provider(filepath string) *File {
	return &File{
		path: filepath,
	}
}

func (f *File) ReadBytes() ([]byte, error) {
	return os.ReadFile(f.path)
}
