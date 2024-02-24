package matrix

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Matrix struct {
	size    [2]int
	vectors []Vector
	storage Vector
}

func New(sizeX, sizeY int) (m *Matrix) {
	m = &Matrix{
		[2]int{sizeX, sizeY},
		make([]Vector, sizeX),
		make(Vector, sizeX*sizeY),
	}

	storage := m.storage
	for i := range m.vectors {
		m.vectors[i], storage = storage[:sizeY], storage[sizeY:]
	}

	return
}

func NewOnes(sizeX, sizeY int) (m *Matrix) {
	m = New(sizeX, sizeY)
	for i := range m.storage {
		m.storage[i] = 1
	}
	return
}

func NewRand(sizeX, sizeY int) (m *Matrix) {
	m = New(sizeX, sizeY)
	for i := range m.storage {
		m.storage[i] = rand.N(byte(2))
	}
	return
}

func NewRange(sizeX, sizeY int) (m *Matrix) {
	m = New(sizeX, sizeY)
	for i := range m.storage {
		m.storage[i] = byte(i)
	}
	return
}

func FromVector(v Vector, sizeX, sizeY int) (m *Matrix, err error) {
	if len(v) != sizeX*sizeY {
		return nil, fmt.Errorf("vector length not equal to sizeX * sizeY")
	}

	m = &Matrix{
		[2]int{sizeX, sizeY},
		make([]Vector, sizeX),
		v,
	}

	for i := range m.vectors {
		m.vectors[i], v = v[:sizeY], v[sizeY:]
	}

	return
}

func (m Matrix) String() string {
	return fmt.Sprint(m.vectors)
}

func (m Matrix) Show() (s string) {
	lines := make([]string, len(m.vectors))
	for i, v := range m.vectors {
		chars := make([]string, len(v))
		for j, x := range v {
			if x == 0 {
				chars[j] = "░░"
			} else {
				chars[j] = "██"
			}
		}
		lines[i] = strings.Join(chars, "")
	}
	return strings.Join(lines, "\n")
}

func (m *Matrix) Shift(direction string) (m1 *Matrix) {
	sizeX, sizeY := m.size[0], m.size[1]
	m1 = New(sizeX, sizeY)
	switch direction {
	case "x+":
		copy(m1.storage[sizeY:], m.storage)
		copy(m1.vectors[0], m.vectors[sizeX-1])
	case "x-":
		copy(m1.storage, m.storage[sizeY:])
		copy(m1.vectors[sizeX-1], m.vectors[0])
	case "y+":
		copy(m1.storage, m.storage)
		for _, v := range m1.vectors {
			last := v[sizeY-1]
			copy(v[1:], v)
			v[0] = last
		}
	case "y-":
		copy(m1.storage, m.storage)
		for _, v := range m1.vectors {
			first := v[0]
			copy(v, v[1:])
			v[sizeY-1] = first
		}
	default:
		panic("")
	}
	return
}

type MatrixSizeMismatchError struct {
	a, b *Matrix
}

func (e MatrixSizeMismatchError) Error() string {
	return fmt.Sprintf("matrix size mismatch: a: %v, b: %v", e.a.size, e.b.size)
}

func (m *Matrix) I(x, y int) byte {
	return m.vectors[x][y]
}

func (m *Matrix) Flip(x, y int) {
	if m.vectors[x][y] == 0 {
		m.vectors[x][y] = 1
	} else {
		m.vectors[x][y] = 0
	}
}

func (m *Matrix) Size(dim int) int {
	return m.size[dim]
}

func (m *Matrix) New() (n *Matrix) {
	return New(m.size[0], m.size[1])
}

func (m *Matrix) Add(n *Matrix) (a *Matrix, err error) {
	if m.size[0] != n.size[0] || m.size[1] != n.size[1] {
		err = MatrixSizeMismatchError{m, n}
		return
	}
	newVec, err := m.storage.Add(n.storage)
	if err != nil {
		return
	}

	a, err = FromVector(newVec, m.size[0], m.size[1])
	return
}

// SumAround sums up the 8 neighbours of each element in a Matrix
func (m *Matrix) SumAround() (s *Matrix) {
	// m1 m1 m2
	// m4 __ m2
	// m4 m3 m3

	m1 := m.Shift("x+")
	m2 := m.Shift("y-")
	m3 := m.Shift("x-")
	m4 := m.Shift("y+")

	s, _ = m1.Add(m2)
	s, _ = s.Add(m3)
	s, _ = s.Add(m4)

	m1 = m1.Shift("y+")
	m2 = m2.Shift("x+")
	m3 = m3.Shift("y-")
	m4 = m4.Shift("x-")

	s, _ = s.Add(m1)
	s, _ = s.Add(m2)
	s, _ = s.Add(m3)
	s, _ = s.Add(m4)

	return
}

func (m *Matrix) Next() (n *Matrix) {
	n = m.New()
	sum := m.SumAround()
	for i, v := range sum.storage {
		switch v {
		case 2:
			n.storage[i] = m.storage[i]
		case 3:
			n.storage[i] = 1
		}
	}
	return
}
