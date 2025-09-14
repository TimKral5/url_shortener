package api

// GetURLResponse is the response after fetching a URL.
type GetURLResponse struct {
	URL string `json:"full_url"`
}

// NewEmptyGetURLResponse constructs an empty GetURLResponse.
func NewEmptyGetURLResponse() *GetURLResponse {
	return &GetURLResponse{
		URL: "",
	}
}

// LoadJSON parses and loads the data from JSON into the current
// object.
func (res *GetURLResponse) LoadJSON(source []byte) error {
	return LoadJSON(res, source)
}

// DumpJSON converts the data from an object into JSON.
func (res *GetURLResponse) DumpJSON() ([]byte, error) {
	return DumpJSON(res)
}
