package model

type Family struct {
	ID           string     `json:"id"`
	LastName     string     `json:"lastName"`
	Parents      []Parents  `json:"parents"`
	Children     []Children `json:"children"`
	Address      Address    `json:"address"`
	CreationDate int64      `json:"creationDate"`
}

type Parents struct {
	FirstName string `json:"firstName"`
}

type Children struct {
	FirstName string `json:"firstName"`
	Gender    string `json:"gender"`
	Grade     int    `json:"grade"`
}

type Address struct {
	State   string
	Country string
	City    string
}

type APIFamily struct {
	ID       string  `json:"id"`
	LastName string  `json:"lastName"`
	Address  Address `json:"address"`
}
