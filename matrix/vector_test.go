package matrix

import "testing"

func TestVectorAddSub(t *testing.T) {
	v1 := make(Vector, 9)
	v2 := make(Vector, 9)

	for i := range 9 {
		v1[i] = byte(i)
		v2[i] = 8 - byte(i)
	}

	a, err := v1.Add(v2)
	if err != nil {
		t.Error(err)
	}

	s, err := v1.Sub(v2)
	if err != nil {
		t.Error(err)
	}
	t.Log(a, s)

	v2 = v2[1:]
	s, err = v1.Add(v2)
	if err != nil {
		t.Log(s == nil, err)
		return
	}
	t.Error("should not get here")
}
