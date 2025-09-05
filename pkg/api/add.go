package api

// AddURLRequest is the request body for creating a URL.
type AddURLRequest struct {
	URL string `json:"full_url"`
}

// AddURLResponse is the response from creating a URL.
type AddURLResponse struct {
	Hash string `json:"hash"`
}


// NewEmptyAddURLRequest constructs an empty AddURLResponse.
func NewEmptyAddURLRequest() *AddURLRequest {
	return &AddURLRequest{
		URL: "",
	}
}

// NewEmptyAddURLResponse constructs an empty AddURLResponse.
func NewEmptyAddURLResponse() *AddURLResponse {
	return &AddURLResponse{
		Hash: "",
	}
}


// LoadJSON parses and loads the data from JSON into the current
// object.
func (res *AddURLRequest) LoadJSON(source []byte) error {
	return LoadJSON(res, source)
}

// DumpJSON converts the data from an object into JSON.
func (res *AddURLRequest) DumpJSON() ([]byte, error) {
	return DumpJSON(res)
}

// LoadJSON parses and loads the data from JSON into the current
// object.
func (res *AddURLResponse) LoadJSON(source []byte) error {
	return LoadJSON(res, source)
}

// DumpJSON converts the data from an object into JSON.
func (res *AddURLResponse) DumpJSON() ([]byte, error) {
	return DumpJSON(res)
}
