package apply_test

type person struct {
	Comment string
	Name    name
}

type name struct {
	First string
	Last  string
}

type book struct {
	Author string
}

var testSet = []book{
	{"Alice"},
	{"Bob"},
	{"Charlie"},
	{"Erin"},
}
