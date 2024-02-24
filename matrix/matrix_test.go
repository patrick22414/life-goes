package matrix

import "testing"

func TestMatrixNew(t *testing.T) {
	m := New(2, 3)
	t.Log(m)
	t.Log(m.storage)

	n := m.New()
	t.Log(n)

	m = NewOnes(2, 3)
	t.Log(m)

	m = NewRand(3, 4)
	t.Log(m)

	m = NewRange(3, 4)
	t.Log(m)

	m1, m2 := m.Shift("x+"), m.Shift("x-")
	t.Log("x+", m1)
	t.Log("x-", m2)

	m1, m2 = m.Shift("y+"), m.Shift("y-")
	t.Log("y+", m1)
	t.Log("y-", m2)

	m1, m2 = NewOnes(3, 3), NewRange(3, 3)
	t.Log(m1.Add(m2))

	m1, m2 = NewOnes(3, 3), NewRange(3, 4)
	t.Log(m1.Add(m2))

	m = NewRand(5, 5)
	t.Logf("\n%v", m.Show())
	t.Log(m.SumAround())

	t.Log("Done")
}
