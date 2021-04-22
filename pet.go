package main

// Pet is the pet struct
type Pet struct {
	ID int64
	Name string
	Category Category
	PhotoUrls []string
	Tags []Tag
	Status string
}