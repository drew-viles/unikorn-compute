// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/unikorn-cloud/core/pkg/openapi"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "https://raw.githubusercontent.com/unikorn-cloud/core/main/pkg/openapi/common.spec.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
