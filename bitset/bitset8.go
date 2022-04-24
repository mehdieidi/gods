package bitset

type Bitset8 uint8

// New returns an empty 8-bit bitset.
func NewBitset8() (b *Bitset8) {
	return
}

// Add gets a position and makes the bit in that position 1. It returns NumberNotFit8Err
// if the given position is larger than 8-bit bitset.
func (b *Bitset8) Add(n uint8) error {
	if n > 7 {
		return NumberNotFit8Err
	}

	*b = *b | 1<<n

	return nil
}

// Containts gets a position and returns true if the bit in that position is 1. It returns NumberNotFit8Err
// if the given position is larger than 8-bit bitset.
func (b *Bitset8) Contains(n uint8) (bool, error) {
	if n > 7 {
		return false, NumberNotFit8Err
	}

	return *b&(1<<n) != 0, nil
}

// Delete gets a position and makes the bit in that position 0. It returns NumberNotFit8Err
// if the given position is larger than 8-bit bitset.
func (b *Bitset8) Delete(n uint8) error {
	if n > 7 {
		return NumberNotFit8Err
	}

	*b = ^(^*b | 1<<n)

	return nil
}

// Intersection gets another 8-bit bitset and returns a bitset containing the intersection of both.
func (b *Bitset8) Intersection(b2 Bitset8) Bitset8 {
	return *b & b2
}

// Union gets another 8-bit bitset and returns a bitset containing the union of both.
func (b *Bitset8) Union(b2 Bitset8) Bitset8 {
	return *b | b2
}

// SymmetricDiff gets another 8-bit bitset and returns a bitset containing the symmetric difference of both.
func (b *Bitset8) SymmetricDiff(b2 Bitset8) Bitset8 {
	return *b ^ b2
}

// Diff gets another 8-bit bitset and returns a bitset containing the difference of both.
func (b *Bitset8) Diff(b2 Bitset8) Bitset8 {
	return *b &^ b2
}

// IncAll increments elements one unit.
func (b *Bitset8) IncAll() {
	*b = *b << 1
}

// DecAll decrements elements one unit.
func (b *Bitset8) DecAll() {
	*b = *b >> 1
}
