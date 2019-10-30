package gen

type swapper interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s swapper) {
	for l := s.Len() - 1; l > 1; l-- {
		s.Swap(mustRandInt(l), l)
	}
	s.Swap(0, mustRandInt(s.Len()))
}
