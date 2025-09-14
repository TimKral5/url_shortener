// Package log provides a suite for logging and diagnostics.
package log

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func logByLevel(prefix string, args ...any) {
	currentTime := time.Now().Format(time.DateTime)

	log.SetFlags(0)
	log.SetPrefix(currentTime + " | " + prefix + ":\033[0m ")
	log.Println(args...)
}

// func Debug(args ...any) {
// 	LogByLevel("\033[34mDBG", args...)
// }

// Log prints a log message to the terminal.
func Log(args ...any) {
	logByLevel("\033[37mLOG", args...)
}

// Info prints a info message to the terminal.
func Info(args ...any) {
	logByLevel("\033[036mINF", args...)
}

// Warn prints a warning to the terminal.
func Warn(args ...any) {
	logByLevel("\033[33mWRN", args...)
}

// Error prints an error message to the terminal.
func Error(args ...any) {
	logByLevel("\033[31mERR", args...)
}

type wrappedWriter struct {
	http.ResponseWriter

	statusCode int
}

// WriteHeader is a wrapper for the original WriteHeader function.
func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

// Middleware is a handler for logging requests to an HTTP server.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, reader *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: writer,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, reader)
		statusPrefix := getHTTPStatusColorPrefix(wrapped.statusCode)
		status := statusPrefix + strconv.Itoa(wrapped.statusCode) + "\033[0m"

		Log(status, reader.Method, reader.URL, "\033[33m"+time.Since(start).String()+"\033[0m")
	})
}

func getHTTPStatusColorPrefix(statusCode int) string {
	switch {
	case statusCode >= 100 && statusCode < 200:
		return "\033[36m"
	case statusCode >= 200 && statusCode < 300:
		return "\033[32m"
	case statusCode >= 300 && statusCode < 400:
		return "\033[35m"
	case statusCode >= 400 && statusCode < 600:
		return "\033[31m"
	default:
		return "\033[33m"
	}
}
