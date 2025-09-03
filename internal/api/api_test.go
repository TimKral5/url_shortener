package api_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	_api "github.com/timkral5/url_shortener/internal/api"
)

var api _api.ShortenerAPI
var hash string

func TestAddURL(t *testing.T) {
	api = _api.NewShortenerAPI()
	mock := httptest.NewUnstartedServer(api.SetupRoutes())

	body := []byte(`{
		"full_url": "https://example.com"
	}`)

	mock.Start()
	defer mock.Close()

	client := mock.Client()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		mock.URL + "/",
		bytes.NewBuffer(body),
	)
	if err != nil {
		t.Error(err)

		return
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)

		return
	}

	var resBody []byte
	_, err = res.Body.Read(resBody)
	if err != nil {
		t.Error("Failed to read the response body.")
		t.Error(err)

		return
	}

	err = res.Body.Close()
	if err != nil {
		t.Error("Response body could not be closed.")
		t.Error(err)

		return
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Status not 200")

		return
	}
}

func TestGetURL(t *testing.T) {
	mock := httptest.NewUnstartedServer(api.SetupRoutes())

	mock.Start()
	defer mock.Close()

	client := mock.Client()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		mock.URL + "/",
		nil,
	)
	if err != nil {
		t.Error(err)

		return
	}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)

		return
	}

	err = res.Body.Close()
	if err != nil {
		t.Error("Response body could not be closed.")

		return
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Status not 200")

		return
	}
}
