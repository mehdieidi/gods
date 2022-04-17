package bitset

import "errors"

type Bitset8 uint8

func New() (b *Bitset8) {
	return
}

func (b *Bitset8) Add(n uint8) error {
	if n > 7 {
		return errors.New("the given number doesn't fit 8-bit set. It must be in range of 0 to 7")
	}

	*b = *b | 1<<n

	return nil
}

func (b *Bitset8) Contains(n uint8) (bool, error) {
	if n > 7 {
		return false, errors.New("the given number doesn't fit 8-bit set. It must be in range of 0 to 7")
	}

	return *b&(1<<n) != 0, nil
}

func (b *Bitset8) Delete(n uint8) error {
	if n > 7 {
		return errors.New("the given number doesn't fit 8-bit set. It must be in range of 0 to 7")
	}

	*b = ^(^*b | 1<<n)

	return nil
}

func (b *Bitset8) Intersection(b2 Bitset8) Bitset8 {
	return *b & b2
}

func (b *Bitset8) Union(b2 Bitset8) Bitset8 {
	return *b | b2
}

func (b *Bitset8) SymmetricDiff(b2 Bitset8) Bitset8 {
	return *b ^ b2
}

func (b *Bitset8) Diff(b2 Bitset8) Bitset8 {
	return *b &^ b2
}

func (b *Bitset8) IncAll() {
	*b = *b << 1
}

func (b *Bitset8) DecAll() {
	*b = *b >> 1
}
