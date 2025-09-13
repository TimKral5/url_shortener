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

// MongoDBInsertError describes an error while inserting a record to
// MongoDB.
type MongoDBInsertError struct {
	InnerError error
	File       string
	Line       int
}

// NewMongoDBInsertError constructs a new MongoDB insert error.
func NewMongoDBInsertError(inner error) *MongoDBConnectError {
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
func (err MongoDBInsertError) Error() string {
	location := err.File + ", " + strconv.Itoa(err.Line) + ": "
	message := "Failed to insert record on MongoDB: " + err.InnerError.Error()

	return location + message
}
