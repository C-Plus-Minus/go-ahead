## What is it?
Fancy functions for fashion-conscious gophers missing their lambdas 🎩🍷

## How to use it?
It's really easy: 

All functions can be found inside the go-ahead/apply package. \
To use one, just "apply" it to your slice and fill out the callback skeleton.
```
index := apply.FindIndex(testkit.Set, func(i int, it testkit.Book) bool {
	return strings.Contains(it.Author, "r")
})
```
For more practical examples take a look at the tests.

## Dependencies?
Nope

