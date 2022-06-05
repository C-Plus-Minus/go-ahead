## What is it?
Fancy functions for fashion-conscious gophers missing their lambdas 🎩🍷

## How to use it?
It's all straight forward: All functions are placed inside the go-ahead/apply package. To use them,
simply "apply" a function to your slice and fill the callback skeleton.
```
// for practical examples take a look at the tests:
index := apply.FindIndex(testkit.Set, func(i int, it testkit.Book) bool {
	return strings.Contains(it.Author, "r")
})
```
