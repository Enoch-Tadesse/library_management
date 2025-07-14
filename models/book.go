package models

const (
	Available = "Available"
	Borrowed  = "Borrowed"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status string
}
