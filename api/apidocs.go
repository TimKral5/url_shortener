package apidocs

import _ "embed"

//go:embed openapi.json
var JSONOpenAPISpecs string
//go:embed openapi.yaml
var YAMLOpenAPISpecs string
