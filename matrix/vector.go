package matrix

import (
	"fmt"
)

type Vector []byte

type VectorSizeMismatchError struct {
	a, b Vector
}

func (e VectorSizeMismatchError) Error() string {
	return fmt.Sprintf("vector size mismatch: a: %v, b: %v", len(e.a), len(e.b))
}

func (v Vector) Add(w Vector) (Vector, error) {
	if len(v) != len(w) {
		return nil, VectorSizeMismatchError{v, w}
	}

	s := make(Vector, len(v))
	for i := range v {
		s[i] = v[i] + w[i]
	}

	return s, nil
}

func (v Vector) Sub(w Vector) (Vector, error) {
	if len(v) != len(w) {
		return nil, VectorSizeMismatchError{v, w}
	}

	s := make(Vector, len(v))
	for i := range v {
		s[i] = v[i] - w[i]
	}

	return s, nil
}
