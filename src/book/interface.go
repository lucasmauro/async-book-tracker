package book

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Read   bool   `json:"read"`
}
