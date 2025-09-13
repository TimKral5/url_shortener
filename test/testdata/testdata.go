// Package testdata provides tools for the generation of test data
// as well as static test cases.
package testdata

import (
	"encoding/json"
	_ "embed"

	"github.com/timkral5/url_shortener/internal/log"
)

//go:embed static.json
var staticData string

type TestData struct {
	TestURLs map[string]string `json:"test_urls"`
}

// ReadStaticTestData parses and returns pre-configured test data.
func ReadStaticTestData() TestData {
	var data TestData

	err := json.Unmarshal([]byte(staticData), &data)
	if err != nil {
		log.Error(err)
		log.Error(staticData)
		return TestData{
			TestURLs: map[string]string{},
		}
	}

	return data
}
