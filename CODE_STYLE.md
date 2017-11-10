# Code Style

This document is only meant to teach you the code style used in this project and will not explain *why* this coding style is used.

* [Tabs vs Spaces](#tabs-vs-spaces)
* [Empty line between blocks and statements](#empty-line-between-blocks-and-statements)
* [Variable names](#variable-names)
* [Types at the top](#types-at-the-top)
* [Private fields at the end of a struct](#private-fields-at-the-end-of-a-struct)
* [Package names](#package-names)
* [Use gofmt](#use-gofmt)
* [Code editor](#code-editor)

## Tabs vs Spaces

Use tabs for indentation and spaces for alignment:

```go
type AnimeTitle struct {
	Romaji    string
	English   string
	Japanese  string
	Canonical string
	Synonyms  []string
}
```

## Empty line between blocks and statements

Bad:

```go
func() {
	if true {
		// Block 1
	}
	if true {
		// Block 2
	}
	doSomething()
	doSomething()
	if true {
		// Block 3
	}
}
```

Good:

```go
func() {
	if true {
		// Block 1
	}

	if true {
		// Block 2
	}

	doSomething()
	doSomething()

	if true {
		// Block 3
	}
}
```

## Variable names

Variables are written in `camelCase` and should clearly state what they contain, preferably with no abbreviations. Common short names like `id` and `url` are allowed.

Iterator variables in loops are an exception to this rule and can be 1-letter, non-significant variable names, usually `i` and `j`. `index` is also quite common.

## Early returns

Do not wrap a whole function in 1 if-block to check parameters. Use early returns.

Bad:

```go
func DoSomething(a string, b string) {
	if a != "" && b != "" {
		doIt(a, b)
	}
}
```

Good:

```go
func DoSomething(a string, b string) {
	if a == "" || b == "" {
		return
	}

	doIt(a, b)
}
```

## Types at the top

`type` definitions should be placed at the very top of your files. Variables, constants, interface implementation assertions and the `package` statement are the only constructs allowed above `type` definitions.

## Private fields at the end of a struct

This is not an absolute rule but try to keep private fields at the end of a struct.

```go
type MyType struct {
	PublicA string
	PublicB string
	PublicC string

	privateA int
}
```

## Package names

Package names should be short lowercase identifiers and tests should be written using the black box pattern. Black box testing can be enabled by adding the suffix `_test` to the package names in `*_test.go` files. It will enable you to test your library like it would be used by another developer, without internal access to private variables.

## Use gofmt

Your IDE should automatically call `gofmt` to format your code every time you save the file.

## Code editor

It is highly recommended to use [Visual Studio Code](https://code.visualstudio.com/) as it has an excellent Go plugin and the repository includes workspace settings to get you started quickly.