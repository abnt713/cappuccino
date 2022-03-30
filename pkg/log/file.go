package log

import (
	"fmt"
	"os"
	"time"
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

// Close closes the file
func (f File) Close() {
	f.Close()
}

// Info writes an info to a file.
func (f File) Info(tag string, info string) {
	now := time.Now()
	fmt.Fprintf(f.file, "%s | info [%s]: %s\n", now.Format(time.RFC3339), tag, info)
}

// Error logs an error.
func (f File) Error(tag string, err error) {
	now := time.Now()
	fmt.Fprintf(f.file, "%s | error [%s]: %s\n", now.Format(time.RFC3339), tag, err.Error())
}
