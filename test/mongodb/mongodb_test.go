// Package main. This package contains several integration tests.
package main

import (
	"os"
	"testing"
	"time"

	"github.com/timkral5/url_shortener/internal/database"
	"github.com/timkral5/url_shortener/test/testdata"
)

const timeout time.Duration = 10 * time.Second

var conn *database.MongoDBConnection

func TestNewMongoDB(t *testing.T) {
	var err error

	connStr := os.Getenv("SHORTENER_MONGODB_URL")

	conn, err = database.NewMongoDBConnection(connStr, timeout)
	if err != nil {
		t.Error("Connection to MongoDB failed:", err)

		return
	}
}

func TestAddURL(t *testing.T) {
	data := testdata.ReadStaticTestData()

	for hash, url := range data.TestURLs {
		hash = hash[:10]

		t.Log("Adding URL", url, "with hash", hash, "...")

		err := conn.AddURL(hash, url)
		if err != nil {
			t.Error("Operation failed:", err)

			return
		}

		t.Log("Operation successful.")
	}
}

func TestGetURL(t *testing.T) {
	data := testdata.ReadStaticTestData()

	for hash, url := range data.TestURLs {
		hash = hash[:10]

		t.Log("Fetching URL for hash", hash, "...")

		result, err := conn.GetURL(hash)
		if err != nil {
			t.Error("Operation failed:", err)

			return
		}

		if result != url {
			t.Error("URL does not match expected value:")
			t.Error("Expected", url, "but got", result)

			return
		}

		t.Log("Operation successful.")
	}
}

func TestDisconnect(t *testing.T) {
	err := conn.Disconnect()
	if err != nil {
		t.Error("Disconnect from MongoDB failed:", err)

		return
	}
}
