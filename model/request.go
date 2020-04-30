package model

type Name struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Joke struct {
	Type  string `json:"type"`
	Value value  `json:"value"`
}

type value struct {
	ID   int    `json:"id"`
	Joke string `json:"joke"`
}
