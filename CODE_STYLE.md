# Code Style

This document is only meant to teach you the code style used in this project and will not explain *why* this coding style is used.

## Tabs vs Spaces

Use **tabs to indent** and **spaces for alignment** only:

```go
type AnimeTitle struct {
	Romaji    string
	English   string
	Japanese  string
	Canonical string
	Synonyms  []string
}
```

## Add an empty line between blocks and other statements

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

Iterator variables in loops are an exception to this rule and can be 1-letter, non-significant variable names, usually `i` and `j`.

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

## Private fields at the end

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