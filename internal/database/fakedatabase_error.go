package database

import (
	"path"
	"runtime"
	"strconv"
)

// FakeConnectError is an emulated connect error.
type FakeConnectError struct {
	File string
	Line int
}

// NewFakeConnectError constructs a new fake connect error.
func NewFakeConnectError() *FakeConnectError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &FakeConnectError{
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err FakeConnectError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to connect to database."

	return location + message
}

// FakeDisconnectError is an emulated disconnect error.
type FakeDisconnectError struct {
	File string
	Line int
}

// NewFakeDisconnectError constructs a new fake disconnect error.
func NewFakeDisconnectError() *FakeDisconnectError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &FakeDisconnectError{
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err FakeDisconnectError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to disconnect from database."

	return location + message
}

// FakeNotFoundError is an emulated disconnect error.
type FakeNotFoundError struct {
	Hash string
	File string
	Line int
}

// NewFakeNotFoundError constructs a new fake disconnect error.
func NewFakeNotFoundError(hash string) *FakeNotFoundError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &FakeNotFoundError{
		Hash: hash,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err FakeNotFoundError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Record not found:" + err.Hash

	return location + message
}
