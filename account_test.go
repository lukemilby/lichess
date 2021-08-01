package lichess

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lukemilby/lichess/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	setUp()

	json := `{"username":"Georges","Profile":{"FirstName": "Bob"  }}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	user, err := client.GetProfile()
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, "Georges", user.Username)
	assert.EqualValues(t, "Bob", user.Profile.FirstName)
}

func TestGetEmail(t *testing.T) {
	setUp()

	json := `{"email":"test@google.com"}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	email, err := client.GetEmail()
	assert.NotNil(t, email)
	assert.Nil(t, err)
	assert.EqualValues(t, "test@google.com", email.Email)
}

func TestSetKidModeStatus(t *testing.T) {
	setUp()
	json := `{"ok":true}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			Body: r,
		}, nil
	}
	status, err := client.SetKidModeStatus(true)
	assert.NotNil(t, status)
	assert.Nil(t, err)
	assert.EqualValues(t, true, status.Ok)
}
