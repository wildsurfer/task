package task

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}

func Test_fetch(t *testing.T) {
	t.Parallel()

	var c http.Client

	endpoint := "/test"
	body := "BODY"

	c.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != endpoint {
			t.Errorf("Wait for '%s' but got '%s'", endpoint, r.URL.Path)
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	})

	f, err := fetch(&c, "/test")

	if err != nil {
		t.Errorf("Error should be nil. Got '%s' instead", err.Error())
	}

	if string(f) != body {
		t.Errorf("Wait for '%s' but got '%s'", body, string(f))
	}

}
