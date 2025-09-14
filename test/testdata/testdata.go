// Package testdata provides tools for the generation of test data
// as well as static test cases.
package testdata

import (
	_ "embed"
	"encoding/json"
	"strconv"

	"github.com/timkral5/url_shortener/internal/hash"
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


// GenerateTestValues makes use of the given seed to generate a single
// set of test values.
func GenerateTestValues(seed int) (string, string) {
	full := hash.GenerateSHA256Hex(strconv.Itoa(seed))
	short := hash.GenerateSHA256Hex(full)

	return short, full
}

// GenerateTestData makes use of the given seed to generate a set
// of testing data.
func GenerateTestData(seed int, size int) TestData {
	urls := map[string]string{}

	for i := range size {	
		short, full := GenerateTestValues(seed + i)
		urls[short] = full
	}

	return TestData{
		TestURLs: urls,
	}
}
