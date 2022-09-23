package gen

import "testing"

const (
	genTestIters  = 5
	genTestLength = 10
)

type intSwapper []int

func TestShuffle(t *testing.T) {
	t.Parallel()

	var (
		input = intSwapper{91, 18, 7, 26, 15, 42, 3}
		prev  = input.CrossDiff()
		cur   int
	)

	for i := 0; i < genTestIters; i++ {
		shuffle(input)

		if cur = input.CrossDiff(); cur == prev {
			t.Fatalf("loop[%d] diff: %d", i, cur)
		}

		prev = cur
	}
}

func TestSimple(t *testing.T) {
	t.Parallel()

	testGenerator(t, Simple)
}

func TestSmart(t *testing.T) {
	t.Parallel()

	testGenerator(t, Smart)
}

func TestSmartRules(t *testing.T) {
	t.Parallel()

	var testRules = rules{
		rule{vow, con},
		rule{vow, vow},
	}

	switch got := testRules.getNext(con); got {
	case vow, con:
	default:
		t.Fatalf("uexpected role: %b", got)
	}
}

func testGenerator(t *testing.T, g Fn) {
	t.Helper()

	seen := make(map[string]struct{})

	for i := 0; i < genTestIters; i++ {
		p := g(genTestLength)

		if _, ok := seen[p]; !ok {
			seen[p] = struct{}{}

			continue
		}

		t.Fatalf("iter[%d] already seen: %s", i, p)
	}
}

func (s intSwapper) Len() int      { return len(s) }
func (s intSwapper) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s intSwapper) CrossDiff() (rv int) {
	for i := 0; i < len(s); i++ {
		if (i & 1) == 1 {
			rv -= s[i]
		} else {
			rv += s[i]
		}
	}

	return rv
}
