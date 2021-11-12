package task

type name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type joke struct {
	Type  string `json:"type"`
	Value struct {
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
		Categories []string `json:"categories"`
	} `json:"value"`
}
