package turing

// infiniteTape is a turing machine tape with "infinite" left and right
type infiniteTape struct {
	buff  []Symbol
	shift int
}

// NewInfiniteTape creates a new infinite tape. An infinite tape is 'infinite'
// in both directions, initialized with the blank symbol (nil) on all positions.
func NewInfiniteTape() Tape {
	buff := make([]Symbol, 0)
	return &infiniteTape{
		buff:  buff,
		shift: 0,
	}
}

func (t *infiniteTape) Get(pos int) (Symbol, error) {
	realPos := pos + t.shift
	if realPos < 0 || realPos >= len(t.buff) {
		return nil, nil
	}
	return t.buff[realPos], nil
}

func (t *infiniteTape) Set(pos int, symbols ...Symbol) error {
	if len(symbols) == 0 {
		return nil
	}

	start := pos + t.shift
	end := start + len(symbols)

	var expandLeft int
	var expandRight int
	if end >= len(t.buff) {
		expandRight = end - len(t.buff)
	}
	if start < 0 {
		expandLeft = -start
		t.shift += -start
		start = 0
	}
	if expandLeft > 0 && expandRight > 0 {
		expandRight += expandLeft
		expandLeft = 0
	}
	if expandRight > 0 {
		t.buff = append(t.buff, make([]Symbol, expandRight)...)
	}
	if expandLeft > 0 {
		t.buff = append(t.buff[:0], append(make([]Symbol, expandLeft), t.buff[0:]...)...)
	}

	copy(t.buff[start:], symbols)
	return nil
}
