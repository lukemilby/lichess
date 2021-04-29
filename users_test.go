package lichess

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lukemilby/lichess/utils/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetUserPublicData(t *testing.T) {
	setUp()
	json := `{"username":"Georges","Profile":{"FirstName": "Bob"  }}`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	user, err := client.GetUserPublicData("some")
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, "Georges", user.Username)
	assert.EqualValues(t, "Bob", user.Profile.FirstName)
}

func TestGetRLUsersStatus(t *testing.T) {
	setUp()

	json := `[{"id":"george","Name":"george", "playing":true},
			{"id":"peter","Name":"peter", "playing":false},
			{"id":"alex","Name":"alex", "playing":true}]
			`

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	user_statuses, err := client.GetRLUsersStatus("george,peter,alex")
	assert.Equal(t, 3, len(*user_statuses))
	assert.Nil(t, err)
	assert.EqualValues(t, "peter", (*user_statuses)[1].Id)
}
