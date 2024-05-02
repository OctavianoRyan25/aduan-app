// internal/storage/storage.go
package storage

import (
	"io"
	"os"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) SaveImage(path string, data io.Reader) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}

	return nil
}

func SaveImage(path string, data io.Reader) error {
	s := NewStorage()
	return s.SaveImage(path, data)
}
