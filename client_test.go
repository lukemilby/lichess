package lichess

import (
	"net/url"
	"testing"

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

// setUp function, add a number to numbers slice
func setUp() {
	client.BaseURL, _ = url.Parse("https://lichess.org")
	client.APIKey = "APIKey"
}
