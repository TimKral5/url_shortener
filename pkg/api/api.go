// Package api provides an interface for communicating with the URL
// Shortener API.
package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

// Error represents an error during any operation regarding the API.
type Error struct {
	Message string
	Inner   error
}

// Client handles and represents all communication to the URL
// Shortener API.
type Client struct {
	connected bool
	baseURL   string
}

// NewEmptyError constructs an empty error with default values.
func NewEmptyError() Error {
	return Error{
		Message: "",
		Inner:   nil,
	}
}

// Error returns the message for this error.
func (err Error) Error() string {
	return err.Message
}

// NewClient constructs a new instance of the APIClient.
func NewClient() Client {
	return Client{
		connected: false,
		baseURL:   "",
	}
}

// Bind sets up a new connection to a given instance of the URL
// shortener.
func (api *Client) Bind(baseURL string) {
	api.baseURL = baseURL
	api.connected = true
}

// Unbind releases a connection set up with Bind().
func (api *Client) Unbind() {
	api.baseURL = ""
	api.connected = false
}

// AddURL can be used to create a new shortened URL.
func (api *Client) AddURL(url string) (*AddURLResponse, error) {
	client := http.DefaultClient

	httpRequest, err := api.setupAddURLRequest(url)
	if err != nil {
		return NewEmptyAddURLResponse(), err
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return NewEmptyAddURLResponse(), Error{
			Message: "Failed to finish the request.",
			Inner:   err,
		}
	}

	if httpResponse.StatusCode != http.StatusOK {
		return NewEmptyAddURLResponse(), Error{
			Message: "The status code is not 200.",
			Inner:   nil,
		}
	}

	response, err := handleAddURLResponse(httpResponse)
	if err != nil {
		return NewEmptyAddURLResponse(), err
	}

	return response, nil
}

// GetURL is used to fetch a full URL from the API.
func (api *Client) GetURL(id string) (*GetURLResponse, error) {
	client := http.DefaultClient

	httpRequest, err := api.setupGetURLRequest(id)
	if err != nil {
		return NewEmptyGetURLResponse(), err
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return NewEmptyGetURLResponse(), Error{
			Message: "Failed to finish the request.",
			Inner:   err,
		}
	}

	if httpResponse.StatusCode != http.StatusOK {
		return NewEmptyGetURLResponse(), Error{
			Message: "The status code is not 200.",
			Inner:   nil,
		}
	}

	response, err := handleGetURLResponse(httpResponse)
	if err != nil {
		return NewEmptyGetURLResponse(), err
	}

	return response, nil
}

func (api *Client) setupAddURLRequest(url string) (*http.Request, error) {
	request := AddURLRequest{
		URL: url,
	}

	requestBytes, err := request.DumpJSON()
	if err != nil {
		return nil, Error{
			Message: "Failed to setup the request body.",
			Inner:   err,
		}
	}

	httpRequest, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		api.baseURL + "/v0/url",
		bytes.NewBuffer(requestBytes),
	)
	if err != nil {
		return nil, Error{
			Message: "Failed to setup the request.",
			Inner:   err,
		}
	}

	return httpRequest, nil
}

func (api *Client) setupGetURLRequest(id string) (*http.Request, error) {
	httpRequest, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		api.baseURL+"/"+id,
		nil,
	)
	if err != nil {
		return nil, Error{
			Message: "Failed to setup the request.",
			Inner:   err,
		}
	}

	return httpRequest, nil
}

func handleAddURLResponse(httpResponse *http.Response) (*AddURLResponse, error) {
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return NewEmptyAddURLResponse(), Error{
			Message: "Failed to read the response body.",
			Inner:   err,
		}
	}

	response := NewEmptyAddURLResponse()

	err = response.LoadJSON(body)
	if err != nil {
		return NewEmptyAddURLResponse(), Error{
			Message: "Failed to parse the response body.",
			Inner:   err,
		}
	}

	err = httpResponse.Body.Close()
	if err != nil {
		return NewEmptyAddURLResponse(), Error{
			Message: "Failed to close the response body.",
			Inner:   err,
		}
	}

	return response, nil
}

func handleGetURLResponse(httpResponse *http.Response) (*GetURLResponse, error) {
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return NewEmptyGetURLResponse(), Error{
			Message: "Failed to read the response body.",
			Inner:   err,
		}
	}

	response := NewEmptyGetURLResponse()

	err = response.LoadJSON(body)
	if err != nil {
		return NewEmptyGetURLResponse(), Error{
			Message: "Failed to parse the response body.",
			Inner:   err,
		}
	}

	err = httpResponse.Body.Close()
	if err != nil {
		return NewEmptyGetURLResponse(), Error{
			Message: "Failed to close the response body.",
			Inner:   err,
		}
	}

	return response, nil
}
