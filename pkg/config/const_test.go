package config

import "testing"

type clampTestCase struct {
	Have, Want int
}

func TestClampCount(t *testing.T) {
	t.Parallel()

	cases := []clampTestCase{
		{MinCount - 1, MinCount},
		{MaxCount + 1, MaxCount},
		{MaxCount / 2, MaxCount / 2},
	}

	for cn, cb := range cases {
		if rv := ClampCount(cb.Have); rv != cb.Want {
			t.Fatalf("case[%d] want: %d got: %d", cn, cb.Want, rv)
		}
	}
}

func TestClampLength(t *testing.T) {
	t.Parallel()

	cases := []clampTestCase{
		{MinLength - 1, MinLength},
		{MaxLength + 1, MaxLength},
		{MaxLength / 2, MaxLength / 2},
	}

	for cn, cb := range cases {
		if rv := ClampLength(cb.Have); rv != cb.Want {
			t.Fatalf("case[%d] want: %d got: %d", cn, cb.Want, rv)
		}
	}
}
