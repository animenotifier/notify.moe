package animediff

// Difference describes a difference between two anime.
type Difference interface {
	Type() string
	Explanation() string
	DetailsA() string
	DetailsB() string
}
