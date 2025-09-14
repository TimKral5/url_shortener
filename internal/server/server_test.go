package server_test

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
	_server "github.com/timkral5/url_shortener/internal/server"
	_api "github.com/timkral5/url_shortener/pkg/api"
	"github.com/timkral5/url_shortener/test/testdata"
)

const testURL string = "https://example.com"
// const testHash string = "100680AD546CE6A577F42F52DF33B4CFDCA756859E664B8D7DE329B150D09CE9"

var server *_server.Server
var mock *httptest.Server

func TestAPIEndpoints(t *testing.T) {
	server = _server.NewServer()

	cache.NewFakeCacheConnection()

	server.Database = database.NewFakeDatabaseConnection()
	server.Cache = cache.NewFakeCacheConnection()

	mock = httptest.NewUnstartedServer(server.SetupRoutes())

	mock.Start()
	defer mock.Close()

	t.Run("TestAddURL", testAddURL)
	t.Run("TestGetURL", testGetURL)
}

func testAddURL(t *testing.T) {
	api := _api.NewClient()
	api.Bind(mock.URL)

	data := testdata.ReadStaticTestData()

	for hash, url := range data.TestURLs {
		hash = strings.ToUpper(hash[:10])

		response, err := api.AddURL(url)
		if err != nil {
			t.Error("Failed to perform the request.")
			t.Error(err)

			apiError := _api.NewEmptyError()
			if errors.As(err, &apiError) {
				t.Error(apiError.Inner)
			}

			return
		}

		if response.Hash != hash {
			t.Error("The received hash does not match the expected value.")
			t.Error("Expected", hash, "but got", response.Hash)
		}
	}
}

func testGetURL(t *testing.T) {
	api := _api.NewClient()
	api.Bind(mock.URL)

	data := testdata.ReadStaticTestData()

	for hash, url := range data.TestURLs {
		hash = strings.ToUpper(hash[:10])

		response, err := api.GetURL(hash)
		if err != nil {
			t.Error("Failed to perform the request.")
			t.Error(err)

			apiError := _api.NewEmptyError()
			if errors.As(err, &apiError) {
				t.Error(apiError.Inner)
			}

			return
		}

		if response.URL != url {
			t.Error("The received URL does not match the expected value.")
			t.Error("Expected", testURL, "but got", response.URL)
		}
	}
}
