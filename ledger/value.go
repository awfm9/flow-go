package ledger

import (
	"bytes"
	"encoding/hex"
)

// Value holds the value part of a ledger key value pair
type Value []byte

// Size returns the value size
func (v Value) Size() int {
	return len(v)
}

func (v Value) String() string {
	return hex.EncodeToString(v)
}

// Equals compares a ledger Value to another one
func (v Value) Equals(other Value) bool {
	if other == nil {
		return false
	}
	return bytes.Equal(v, other)
}