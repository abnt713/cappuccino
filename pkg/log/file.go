package log

import (
	"fmt"
	"os"
)

// NewFile creates a new file logger instance.
func NewFile(filepath string) (File, error) {
	var file *os.File
	if _, err := os.Stat(filepath); err == nil {
		file, err = os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			file.Close()
			return File{}, err
		}
	} else {
		file, err = os.Create(filepath)
		if err != nil {
			file.Close()
			return File{}, err
		}
	}

	return File{file: file}, nil
}

// File logs events to a file.
type File struct {
	file *os.File
}

// Error logs an error.
func (f File) Error(tag string, err error) {
	fmt.Fprintf(f.file, "error [%s]: %s\n", tag, err.Error())
}
