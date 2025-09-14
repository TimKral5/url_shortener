package database

import (
	"path"
	"runtime"
	"strconv"
)

// MongoDBConnectError describes an error while connecting to
// MongoDB.
type MongoDBConnectError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MongoDBConnectError)(nil)

// NewMongoDBConnectError constructs a new MongoDB connect error.
func NewMongoDBConnectError(inner error) *MongoDBConnectError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MongoDBConnectError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MongoDBConnectError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to connect to MongoDB: " + err.InnerError.Error()

	return location + message
}

// MongoDBDisconnectError describes an error while connecting to
// MongoDB.
type MongoDBDisconnectError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MongoDBDisconnectError)(nil)

// NewMongoDBDisconnectError constructs a new MongoDB connect error.
func NewMongoDBDisconnectError(inner error) *MongoDBDisconnectError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MongoDBDisconnectError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MongoDBDisconnectError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to disconnect from MongoDB: " + err.InnerError.Error()

	return location + message
}

// MongoDBInsertError describes an error while inserting a record to
// MongoDB.
type MongoDBInsertError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MongoDBInsertError)(nil)

// NewMongoDBInsertError constructs a new MongoDB insert error.
func NewMongoDBInsertError(inner error) *MongoDBInsertError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MongoDBInsertError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MongoDBInsertError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to insert record on MongoDB: " + err.InnerError.Error()

	return location + message
}

// MongoDBFetchError describes an error while fetching a record from
// MongoDB.
type MongoDBFetchError struct {
	InnerError error
	File       string
	Line       int
}

var _ error = (*MongoDBFetchError)(nil)

// NewMongoDBFetchError constructs a new MongoDB fetch error.
func NewMongoDBFetchError(inner error) *MongoDBFetchError {
	_, file, line, _ := runtime.Caller(1)
	root := path.Join(file, "../../../")
	file = file[len(root):]

	return &MongoDBFetchError{
		InnerError: inner,
		File:       file,
		Line:       line,
	}
}

// Error returns the error's message.
func (err MongoDBFetchError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to fetch record from MongoDB: " + err.InnerError.Error()

	return location + message
}
