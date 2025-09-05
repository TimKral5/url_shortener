package server_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	_server "github.com/timkral5/url_shortener/internal/server"
	_api "github.com/timkral5/url_shortener/pkg/api"
)

const testURL string = "https://example.com"
const testHash string = "100680AD546CE6A577F42F52DF33B4CFDCA756859E664B8D7DE329B150D09CE9"

var server _server.Server
var mock *httptest.Server

var hash string


func TestAPIEndpoints(t *testing.T) {
	server = _server.NewServer()
	mock = httptest.NewUnstartedServer(server.SetupRoutes())

	mock.Start()
	defer mock.Close()

	t.Run("TestAddURL", testAddURL)
	t.Run("TestGetURL", testGetURL)
}

func testAddURL(t *testing.T) {	
	api := _api.NewClient()
	api.Bind(mock.URL)

	response, err := api.AddURL(testURL)
	if err != nil {
		t.Error("Failed to perform the request.")
		t.Error(err)

		apiError := _api.NewEmptyError()
		if errors.As(err, &apiError) {
			t.Error(apiError.Inner)
		}

		return
	}

	if response.Hash != testHash[:10] {
		t.Error("The received hash does not match the expected value.")
		t.Error("Expected", testHash[:10], "but got", response.Hash)
	}

	hash = response.Hash
}

func testGetURL(t *testing.T) {
	api := _api.NewClient()
	api.Bind(mock.URL)

	response, err := api.GetURL(testHash[:10])
	if err != nil {
		t.Error("Failed to perform the request.")
		t.Error(err)

		apiError := _api.NewEmptyError()
		if errors.As(err, &apiError) {
			t.Error(apiError.Inner)
		}

		return
	}

	if response.URL != testURL {
		t.Error("The received URL does not match the expected value.")
		t.Error("Expected", testURL, "but got", response.URL)
	}
}
