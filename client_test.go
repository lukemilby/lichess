package lichess

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/lukemilby/lichess/utils/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	client = new(Client)
)

func TestNewReuqestWithoutUrl(t *testing.T) {
	setUp()
	client.BaseURL = nil
	_, err := client.newRequest("POST", "/get_some", nil)
	assert.Error(t, err)
}

func TestNewReuqestWithoutApiKey(t *testing.T) {
	setUp()
	client.APIKey = ""
	_, err := client.newRequest("POST", "/get_some", nil)
	assert.Error(t, err)
}

func TestNewReuqestSuccess(t *testing.T) {
	setUp()
	req, err := client.newRequest("POST", "/get_some", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)
}

func TestDo(t *testing.T) {
	setUp()
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	req, err := client.newRequest("GET", "/some", nil)
	resp, err := client.HttpClient.Do(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, 200, resp.StatusCode)
	assert.EqualValues(t, ioutil.NopCloser(bytes.NewReader([]byte(json))), resp.Body)
}

// setUp function, add a number to numbers slice
func setUp() {
	client.BaseURL, _ = url.Parse("https://lichess.org")
	client.APIKey = "APIKey"
	client.HttpClient = &mocks.MockClient{}
}
