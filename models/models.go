package models


// Book struct (model)
type Book struct {
	ID		string `json: "id"`
	Isbn	string `json: "isbn"`
	Title	string `json: "title"`
	Author	*Author `json: "author"`
}

// Library book slice
var Library = []Book{
	Book{
		ID: "1",
		Title: "Kubwa Kiti",
		Isbn: "65431",
		Author: &Author {
			ID: "1",
			Firstname: "Florian",
			Lastname: "GOTO",
		},
	},
}

// Author struct (model)
type Author struct {
	ID		string `json: "id"`
	Firstname	string `json: "firstname"`
	Lastname	string `json: "lastname"`
}
