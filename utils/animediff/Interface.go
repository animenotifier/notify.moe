package animediff

// Difference describes a difference between two anime.
type Difference interface {
	TypeName() string
	Explanation() string
	DetailsA() string
	DetailsB() string
	Hash() uint64
}
