package task

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
