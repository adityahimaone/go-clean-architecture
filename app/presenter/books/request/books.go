package request

type Book struct {
	Title  string `json:"title"`
	ISBN   string `json:"isbn"`
	Writer string `json:"writer"`
	Status bool   `json:"status"`
}
