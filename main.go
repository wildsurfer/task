package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Joke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}

var httpClient http.Client

func init() {
	httpClient = http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":5000", RequestLogger(mux)))
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func httpHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/text")

	name, errName := fetchRandomName()
	joke, errJoke := fetchJoke(name)

	if errName != nil || errJoke != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Sorry, something went wrong. Please try again later..."))
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(joke.Value.Joke + "\n"))
	}
}

func fetchRandomName() (name Name, err error) {
	const endpoint = "https://names.mcquay.me/api/v0/"
	req, reqErr := http.NewRequestWithContext(context.Background(), http.MethodGet, endpoint, nil)

	if reqErr != nil {
		err = reqErr
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		err = getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		err = readErr
	}

	jsonErr := json.Unmarshal(body, &name)
	if jsonErr != nil {
		err = jsonErr
	}

	return name, err
}

func fetchJoke(name Name) (joke Joke, err error) {
	endpoint := fmt.Sprintf(
		"http://api.icndb.com/jokes/random?firstName=%s&lastName=%s",
		name.FirstName,
		name.LastName,
	)
	req, reqErr := http.NewRequestWithContext(context.Background(), http.MethodGet, endpoint, nil)

	if reqErr != nil {
		err = reqErr
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		err = getErr
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		err = readErr
	}

	jsonErr := json.Unmarshal(body, &joke)
	if jsonErr != nil {
		err = jsonErr
	}

	return joke, err
}
