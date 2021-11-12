package task

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func httpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 2 * time.Second,
	}

	return client
}

func requestLogger(targetMux http.Handler) http.Handler {
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

	cli := httpClient()
	name, errName := fetchRandomName(cli)
	joke, errJoke := fetchJoke(cli, name)

	if errName != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Sorry, something went wrong.\n" +
			"Can't get random name from remote API. \n" +
			"Please try again later...\n\n" +
			errName.Error() + "\n"))
	} else if errJoke != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Sorry, something went wrong.\n" +
			"Can't get the joke from remote API.\n" +
			"Please try again later...\n\n" +
			errJoke.Error() + "\n"))
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(joke.Value.Joke + "\n"))
	}
}

//ListenAndServe starts the server on port 5000
func ListenAndServe() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpHandler)

	logger := requestLogger(mux)
	err := http.ListenAndServe(":5000", logger)

	return fmt.Errorf("server error: %w", err)
}
