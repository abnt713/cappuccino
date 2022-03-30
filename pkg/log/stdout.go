package log

import "fmt"

// NewSTDOut creates a new STDOut logger instance.
func NewSTDOut() STDOut {
	return STDOut{}
}

// STDOut logs events using stdout.
type STDOut struct{}

// Error logs an error.
func (STDOut) Error(tag string, err error) {
	fmt.Printf("error [%s]: %s", tag, err.Error())
}
