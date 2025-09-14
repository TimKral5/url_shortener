package cache

import (
	"path"
	"runtime"
	"strconv"
)

// MemcachedDisconnectError describes an error during disconnect.
type MemcachedDisconnectError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MemcachedDisconnectError)(nil)

// NewMemcachedDisconnectError constructs a new Memcached disconnect error.
func NewMemcachedDisconnectError(inner error) *MemcachedDisconnectError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MemcachedDisconnectError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MemcachedDisconnectError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to disconnect from Memcached: " + err.InnerError.Error()

	return location + message
}

// MemcachedAddError describes an error while adding a new record to
// Memcached.
type MemcachedAddError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MemcachedAddError)(nil)

// NewMemcachedAddError constructs a new Memcached set error.
func NewMemcachedAddError(inner error) *MemcachedAddError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MemcachedAddError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MemcachedAddError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to add record to Memcached: " + err.InnerError.Error()

	return location + message
}

// MemcachedGetError describes an error while fetching a record from
// Memcached.
type MemcachedGetError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MemcachedGetError)(nil)

// NewMemcachedGetError constructs a new Memcached fetch error.
func NewMemcachedGetError(inner error) *MemcachedGetError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MemcachedGetError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MemcachedGetError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to fetch record from Memcached: " + err.InnerError.Error()

	return location + message
}
