package db

// CardComplete is a complete card
type CardComplete struct {
	Name       string
	Amount     int
	Comment    string
	Colors     []int
	Types      []int
	Subtypes   []int
	Supertypes []int
}
