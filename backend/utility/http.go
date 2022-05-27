package utility

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// Get Request ...
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("new request failed")
	}
	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	client := &http.Client{}
	return client.Do(req)
}

// PostJSON request ...
func PostJSON(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	// add body
	var bodyJSON []byte
	var req *http.NewRequest
	if body != nil {
		var err error
		bodyJSON, err = json.Marshal(body)
		if err != nil {
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, errors.New("new request failed")
	}
	req.Header.Add("Content-type", "application/json")
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	req.URL.RawQuery = q.Encode()
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	client := &http.Client{}
	return client.Do(req)
}

// PostFormURLEnCode request ...
func PostFormURLEnCode() {
	
}
