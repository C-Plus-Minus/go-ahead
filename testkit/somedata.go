package testkit

type Person struct {
	Comment string
	Name    Name
}

type Name struct {
	First string
	Last  string
}

type Book struct {
	Author string
}

var Set = []Book{
	{"Alice"},
	{"Bob"},
	{"Charlie"},
	{"Erin"},
}
