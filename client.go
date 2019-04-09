package lichess

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL *url.URL
	UserAgent string
	APIKey string
	HttpClient *http.Client
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error){
	rel := &url.URL{Path:path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	//if c.APIKey == "" {
	//	log.Fatal("Missing API key")
	//}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	return req, nil
}

func (c *Client) do(req *http.Request,
	v interface{}) (response *http.Response, err error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return response, err
}
