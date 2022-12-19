package storage

import "os"

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(s string) error
}

type Storage interface {
	Reader
	Writer
}

// fileStorage is a basic implementation for read and writing files. The struct count with one file tha is going to be
// read (filenameIn - README.md.tpl) and is going to be written in another file out (filenameOut - README.md).
type fileStorage struct {
	filenameIn, filenameOut string
}

func New(filenameIn, filenameOut string) Storage {
	return fileStorage{
		filenameIn:  filenameIn,
		filenameOut: filenameOut,
	}
}

func (f fileStorage) Read() (string, error) {
	data, err := os.ReadFile(f.filenameIn)
	if err != nil {
		return "", nil
	}

	return string(data), nil
}

func (f fileStorage) Write(s string) error {
	return os.WriteFile(f.filenameOut, []byte(s), 0644)
}
