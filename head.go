package turing

import (
	"fmt"
	"strings"
)

// Head is a reading head
type Head struct {
	tape   Tape
	pos    int
	minPos int
	maxPos int
}

// Move moves the head or stays at the same place
func (h *Head) Move(movement string) {
	switch movement {
	case LEFT:
		h.pos--
		if h.pos < h.minPos {
			h.minPos = h.pos
		}
	case RIGHT:
		h.pos++
		if h.pos > h.maxPos {
			h.maxPos = h.pos
		}
	default:
	}
}

// Attach attachs the head to a tape at the specified position.
func (h *Head) Attach(tape Tape, pos int) {
	h.pos = pos
	h.tape = tape
	h.minPos = pos
	h.maxPos = pos
}

// Read reads the current Symbol under the Head
func (h *Head) Read() (Symbol, error) {
	return h.tape.Get(h.pos)
}

// Write writes a Symbol under the Head
func (h *Head) Write(s Symbol) error {
	return h.tape.Set(h.pos, s)
}

// Pos returns the head position on the attached tape.
func (h *Head) Pos() int {
	return h.pos
}

// MinPos returns the smallest position that the head after being attached
// to the tape.
func (h *Head) MinPos() int {
	return h.minPos
}

// MaxPos returns the biggest position that the head after being attached
// to the tape.
func (h *Head) MaxPos() int {
	return h.maxPos
}

// PrintTape prints the tape on all positions that head walked.
func (h *Head) PrintTape(from, to int) string {
	builder := strings.Builder{}
	for i := from; i <= to; i++ {
		v, _ := h.tape.Get(i)
		line := fmt.Sprintf(" %d: %v\n", i, v)
		if h.pos == i {
			line = fmt.Sprintf("[%d: %v]\n", i, v)
		}
		builder.WriteString(line)
	}
	return builder.String()
}
