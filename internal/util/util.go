package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func BuildRequest(method string, baseURL string, path string, pathArgs map[string]interface{}, body interface{}) (*http.Request, error) {

	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, baseURL+path, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	v := url.Values{}

	for key, value := range pathArgs {
		v.Add(key, fmt.Sprintf("%v", value))
	}

	req.URL.RawQuery = v.Encode()

	return req, nil
}
