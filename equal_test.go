package deep

import (
	"math"
	"testing"
)

var m1 = map[int]any{1: map[int]int{1: 1}, 2: []any{4}, 4: 8, 8: 16}

func TestEqualMaps(t *testing.T) {
	if !EqualMaps(m1, m1) {
		t.Errorf("Equal(%v, %v) = false, want true", m1, m1)
	}
	if EqualMaps(m1, (map[int]any)(nil)) {
		t.Errorf("Equal(%v, nil) = true, want false", m1)
	}
	if EqualMaps((map[int]any)(nil), m1) {
		t.Errorf("Equal(nil, %v) = true, want false", m1)
	}
	if !EqualMaps[map[int]any, map[int]any](nil, nil) {
		t.Error("Equal(nil, nil) = false, want true")
	}
	if ms := map[int]any{1: 2}; EqualMaps(m1, ms) {
		t.Errorf("Equal(%v, %v) = true, want false", m1, ms)
	}

	// Comparing NaN for equality is expected to fail.
	mf := map[int]float64{1: 0, 2: math.NaN()}
	if EqualMaps(mf, mf) {
		t.Errorf("Equal(%v, %v) = true, want false", mf, mf)
	}
}

var equalIntTests = []struct {
	s1, s2 []int
	want   bool
}{
	{
		[]int{1},
		nil,
		false,
	},
	{
		[]int{},
		nil,
		true,
	},
	{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		true,
	},
	{
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
		false,
	},
}

var equalFloatTests = []struct {
	s1, s2       []float64
	wantEqual    bool
	wantEqualNaN bool
}{
	{
		[]float64{1, 2},
		[]float64{1, 2},
		true,
		true,
	},
	{
		[]float64{1, 2, math.NaN()},
		[]float64{1, 2, math.NaN()},
		false,
		true,
	},
}

func TestEqualSlices(t *testing.T) {
	for _, test := range equalIntTests {
		if got := EqualSlices(test.s1, test.s2); got != test.want {
			t.Errorf("Equal(%v, %v) = %t, want %t", test.s1, test.s2, got, test.want)
		}
	}
	for _, test := range equalFloatTests {
		if got := EqualSlices(test.s1, test.s2); got != test.wantEqual {
			t.Errorf("Equal(%v, %v) = %t, want %t", test.s1, test.s2, got, test.wantEqual)
		}
	}
}
