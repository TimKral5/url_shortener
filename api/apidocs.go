// Package apidocs embeds the API documentation files into the
// binary.
package apidocs

import _ "embed"

// JSONOpenAPISpecs is the JSON OpenAPI specification file.
//go:embed openapi.json
var JSONOpenAPISpecs string

// YAMLOpenAPISpecs is the YAML OpenAPI specification file.
//go:embed openapi.yaml
var YAMLOpenAPISpecs string

// HTMLOpenAPIDocs is the HTML OpenAPI documentation file.
//go:embed build/html/index.html
var HTMLOpenAPIDocs string

// HTMLOpenAPIRefs is the HTML OpenAPI reference file.
//go:embed build/html2/index.html
var HTMLOpenAPIRefs string
