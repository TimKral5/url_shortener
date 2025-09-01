package cache_test

import (
	"testing"

	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/hash"
)

func TestConnect(t *testing.T) {
	t.Parallel()

	mock := cache.NewFakeCacheConnection()

	var conn cache.Connection = mock

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

func TestAddURL(t *testing.T) {
	t.Parallel()

	ctrlURL := "https://example.com"
	ctrlHash := "100680AD54"

	var conn cache.Connection = cache.NewFakeCacheConnection()

	hash := hash.GenerateSHA256Hex(ctrlURL)[:10]

	err := conn.AddURL(hash, ctrlURL)
	if err != nil {
		t.Error(err)

		return
	}

	if hash != ctrlHash {
		t.Error("Hash does not match the expected value.")

		return
	}

	full, err := conn.GetURL(hash)
	if err != nil {
		t.Error(err)

		return
	}

	if full != ctrlURL {
		t.Error("The full URL does not match the expected value.")
	}
}
