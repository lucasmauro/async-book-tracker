package book

type Book struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Read   bool   `json:"read,omitempty"`
}
