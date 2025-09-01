package database_test

import (
	"testing"

	"github.com/timkral5/url_shortener/internal/database"
)

func TestConnect(t *testing.T) {
	t.Parallel()

	mock := database.NewFakeDatabaseConnection()

	var conn database.Connection = mock

	err := conn.Connect("")
	if err != nil {
		t.Error(err)

		return
	}

	err = conn.Disconnect()
	if err != nil {
		t.Error(err)

		return
	}

	mock.FailConnect = true
	mock.FailDisconnect = true

	err = conn.Connect("")
	if err == nil {
		t.Error("There was no error while connecting.")
	}

	err = conn.Disconnect()
	if err == nil {
		t.Error("There was no error while disconnecting.")
	}
}

func TestCreateURL(t *testing.T) {
	t.Parallel()

	url := "https://example.com"
	urlHash := "100680AD54"

	var conn database.Connection = database.NewFakeDatabaseConnection()

	hash, err := conn.CreateURL(url)
	if err != nil {
		t.Error(err)

		return
	}

	if hash != urlHash {
		t.Error("Hash does not match the expected value.")

		return
	}

	full, err := conn.GetURL(hash)
	if err != nil {
		t.Error(err)

		return
	}

	if full != url {
		t.Error("The full URL does not match the expected value.")
	}
}
