// Package view provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package view

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RYwW7bRhD9FWLbIy1RdpIGAnKI66IQULRGcjR8YMS1zJTi0rtLw4JBwLLRNoWBBkVz",
	"KAqkRpsfkNUKVu1K/oXZPypmSUmkSMm2IiPpxTCXy5k382Zm3+qQ1FkzYD71pSDVQ8KpCJgvqH74gnPG",
	"nyUruOBQUeduIF3mkyqBM3UEQ9VWx9BRrw0YqlfQh3O4hA4xScBZQLl0Y1NNKoTdKDLyJ/TgEo2kDfSJ",
	"SeiB3Qw8OnePbAW4QUju+g0SRZFJnjO/IdZbNScNvM58SX2J/9pB4Ll1G92XXwrEcEhEfZc2bf3W877Z",
	"IdWtQ/IppzukSj4pTxJUjveJctpH6EkSbWvXU5G9U224hp56BQN1Cv8g9mPoQk8dkxHOr1wh7xtn4mMx",
	"nJGZONUsbjIh0eYzuhdSoWFmaW5wFgYFJL+FDlxAFzowgD70DPhLHakTuIZrdZql+pfRC11EWX5NcrDC",
	"7MBdqTOHNqi/Qg8kt1ek3dDOX7i+g9uqhNO90OXUifEzXLsZEqagjQ9ZQGfxsnq9FDiROXmqbiX5SjBu",
	"jx2wFy9pXWKVbIb/i4SzpitpM5AtnXHP9b8twPSHaqtTuML5YMAA/yQpVz9lIc3fuBSInHrUFnTDlkUz",
	"6Q10cKgZ6gh6cAV9uNDJSUGcsWMp2D58webx1DYKEP0Kb2aiyL1LoLi+pA3KF+xlSQ9kAZDfdAVf4chS",
	"pzMhzdm1hDxNdXaStKKefp7wO6OTb9mbkUlcJ7O/srr2IJfnVDsu2GH5bpkYqjwqWZ+VVi3rUdFno0JO",
	"OZ5bxDkDI7oXIXGaENchZnbeZuNKvCXpmsVb+tQvmHBwDX3VhiFcjqD1YJATQ44t9SF+09Gdi0F/uD0l",
	"HpYCBetY3A7TODM253ZrFkZcdv0dhjalKzV5GrXxdLNGTLJPuYihVkpWyUKzLKC+HbikStZKVmkNsdpy",
	"V6Mq24Fb3q+URxUVMCEL5/YQzqEDXSyQUYUNYAhdGKKwSVcLZkGrq5pDqmNlQ+JgqJDrzGndSZLNy9y0",
	"cIqSeZES26uWNcvKeF85L261nNuxkyKY/3VWz+vRHjabNm/dKXcGdA041yIcDz4Ui6jOT7CD9ITcSp9B",
	"WK1BWETW73Ce2E85/FtPhoE6Vd/NpSu8V7bCj52sO+SumC7ozaIrMjPdVj50nShmz6OFYumddpzCkXLd",
	"h4si53lCN7R1zNh6q7ahW5/bTSopF/qWk/VZ28gWh4uLOCyISXy7qc8/h6TnkuQhNVP0T5+S0XaO4Qf5",
	"WL9mxudJgS2NykXSZ6gfoKel54WBSmt26zVoUeudwRCu1AlaeR/SvqTyo2Lsw/bkYkm9LZVTbRlLxtuS",
	"W+QXuigJJqMiVgc5gp96ns7ZTQTDWxjCvxiKodrqWB3FEk99r++bmu69kPLWhO/AbqDqmjA8pqFSpGLn",
	"SvIdxpuJSTMx86SCBOWvCpgd6Oust/XvHEMEfKJ+xIGql3p43KEozgaiR2ZhIJRvzo7l4Z2DmdwvPEmf",
	"VCzLIuY4wtjXJMiH79UJmd+f7rMTbl+B4+r/OfdJJ70zlpiC8v3igvRY3fYM6u8Tk4TcI1WyK2VQLZf1",
	"i10mZPWx9dgi0Xb0XwAAAP//sqsxGgUVAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
