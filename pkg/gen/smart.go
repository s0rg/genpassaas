package gen

type role byte

const (
	vow role = 1
	con role = 2
)

var defaultRoles = []role{vow, con}

func randomRole() role {
	return defaultRoles[mustRandInt(len(defaultRoles))]
}

type rule struct {
	l, r role
}

type rules []rule

var defaultRules = rules{
	rule{vow, con},
	rule{con, vow},
	rule{vow, vow},
}

func (rls rules) getNext(r role) role {
	var (
		match []*rule
		idx   int
	)

	for i := 0; i < len(rls); i++ {
		if rul := &rls[i]; rul.l == r {
			match = append(match, rul)
		}
	}

	switch c := len(match); c {
	case 0:
		return randomRole()
	case 1:
		// idx still 0
	default:
		idx = mustRandInt(c)
	}

	return match[idx].r
}

type char struct {
	r role
	c []byte
}

func (c char) getByte() byte {
	return c.c[mustRandInt(len(c.c))]
}

var defaultAlphabet = []char{
	{vow, []byte("aA4@")},
	{con, []byte("bB8")},
	{con, []byte("cC(")},
	{con, []byte("dD)")},
	{vow, []byte("eE3")},
	{con, []byte("fF")},
	{con, []byte("gG6")},
	{con, []byte("hH#")},
	{vow, []byte("iI!")},
	{con, []byte("jJ]")},
	{con, []byte("kK")},
	{con, []byte("lL1|")},
	{con, []byte("mM")},
	{con, []byte("nN")},
	{vow, []byte("oO0")},
	{con, []byte("pP9")},
	{con, []byte("qQ")},
	{con, []byte("rR")},
	{con, []byte("sS$5")},
	{con, []byte("tT7")},
	{vow, []byte("uU")},
	{con, []byte("vV")},
	{con, []byte("wW")},
	{con, []byte("xX")},
	{vow, []byte("yY")},
	{con, []byte("zZ2")},
}

type chars []*char

func (c chars) Len() int {
	return len(c)
}

func (c chars) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type charBuf map[role]chars

func (cb charBuf) addChar(c *char) {
	cb[c.r] = append(cb[c.r], c)
}

func (cb charBuf) getByte(r role) byte {
	chrs := cb[r]

	return chrs[mustRandInt(len(chrs))].getByte()
}

func (cb charBuf) shuffle() {
	for _, chrs := range cb {
		shuffle(chrs)
	}
}

func buildCharBuf() charBuf {
	cb := make(charBuf)

	for _, ch := range defaultAlphabet {
		cb.addChar(&char{ch.r, ch.c})
	}

	cb.shuffle()

	return cb
}

func Smart(length int) string {
	var (
		pass    []byte
		chars   = buildCharBuf()
		rules   = defaultRules
		curRole = randomRole()
	)

	for len(pass) < length {
		pass = append(pass, chars.getByte(curRole))
		curRole = rules.getNext(curRole)
	}

	return string(pass)
}
