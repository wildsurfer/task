package task

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchRandomName(cli *http.Client) (name name, err error) {
	const endpoint = "https://names.mcquay.me/api/v0/"

	body, fetchErr := fetch(cli, endpoint)
	if fetchErr != nil {
		err = fmt.Errorf("random name error: %w", fetchErr)

		return
	}

	jsonErr := json.Unmarshal(body, &name)
	if jsonErr != nil {
		err = fmt.Errorf("unmarshal error: %w", jsonErr)
	}

	return name, err
}

func fetchJoke(cli *http.Client, name name) (joke joke, err error) {
	endpoint := fmt.Sprintf(
		"http://api.icndb.com/jokes/random?firstName=%s&lastName=%s",
		name.FirstName,
		name.LastName,
	)

	body, fetchErr := fetch(cli, endpoint)
	if fetchErr != nil {
		err = fmt.Errorf("joke error: %w", fetchErr)

		return
	}

	jsonErr := json.Unmarshal(body, &joke)
	if jsonErr != nil {
		err = jsonErr
	}

	return joke, err
}

func fetch(httpClient *http.Client, endpoint string) (body []byte, err error) {
	req, reqErr := http.NewRequestWithContext(context.Background(), http.MethodGet, endpoint, nil)

	if reqErr != nil {
		err = fmt.Errorf("request error: %w", reqErr)

		return
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		err = fmt.Errorf("response error: %w", getErr)

		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		err = fmt.Errorf("read body error: %w", readErr)
	}

	return
}
