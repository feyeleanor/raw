package raw

import "testing"

func TestMerge(t *testing.T) {
	SHOULD_MATCH := "Map elements m[%v] and n[%v] should match but are %v and %v"

	m := NewMap(map[int]int{ 0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9 })
	n := NewMap(map[int]int{ 10: 0, 11: 1, 12: 2, 13: 3, 14: 4, 15: 5, 16: 6, 17: 7, 18: 8, 19: 9 })
	Merge(m, n)

	switch {
	case m.Len() == n.Len():	t.Fatalf("Map length should be %v not %v", 20, m.Len())
	case m.At(10) != n.At(10):	t.Fatalf(SHOULD_MATCH, 10, 10, m.At(10), n.At(10))
	case m.At(11) != n.At(11):	t.Fatalf(SHOULD_MATCH, 11, 11, m.At(11), n.At(11))
	case m.At(12) != n.At(12):	t.Fatalf(SHOULD_MATCH, 12, 12, m.At(12), n.At(12))
	case m.At(13) != n.At(13):	t.Fatalf(SHOULD_MATCH, 13, 13, m.At(13), n.At(13))
	case m.At(14) != n.At(14):	t.Fatalf(SHOULD_MATCH, 14, 14, m.At(14), n.At(14))
	case m.At(15) != n.At(15):	t.Fatalf(SHOULD_MATCH, 15, 15, m.At(15), n.At(15))
	case m.At(16) != n.At(16):	t.Fatalf(SHOULD_MATCH, 16, 16, m.At(16), n.At(16))
	case m.At(17) != n.At(17):	t.Fatalf(SHOULD_MATCH, 17, 17, m.At(17), n.At(17))
	case m.At(18) != n.At(18):	t.Fatalf(SHOULD_MATCH, 18, 18, m.At(18), n.At(18))
	case m.At(19) != n.At(19):	t.Fatalf(SHOULD_MATCH, 19, 19, m.At(19), n.At(19))
	}
}