package animediff

// Difference describes a difference between two anime.
type Difference interface {
	String() string
	DetailsA() string
	DetailsB() string
}
