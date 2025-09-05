package api

import "encoding/json"

// LoadJSON parses and inserts data from JSON.
func LoadJSON(dest any, source []byte) error {
	err := json.Unmarshal(source, dest)
	if err != nil {
		return Error{
			Message: "Failed to parse the JSON.",
			Inner: err,
		}
	}

	return nil
}

// DumpJSON converts the data from an object into JSON.
func DumpJSON(obj any) ([]byte, error) {
	buf, err := json.Marshal(obj)
	if err != nil {
		return nil, Error{
			Message: "Failed to dump the JSON.",
			Inner: err,
		}
	}

	return buf, nil
}
