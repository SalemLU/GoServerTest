package main

// Users struct which contains
// an array of users
type Movies struct {
	Movies []Movie `json:"movies"`
}

// User struct which contains a name
// a type and a list of social links
type Movie struct {
	Title    string   `json:"title"`
	Year     int64    `json:"year"`
	Runtime  int64    `json:"runtime"`
	Genres   []string `json:"genres"`
	Director string   `json:"director"`
	Actors   string   `json:"actors"`
}
