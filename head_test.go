package turing_test

import (
	"testing"

	"github.com/massahud/turing"
	"github.com/stretchr/testify/assert"
)

type MockTape struct {
	setPos     int
	getPos     int
	setSymbols []turing.Symbol
}

func (m *MockTape) Get(pos int) (turing.Symbol, error) {
	m.getPos = pos
	return nil, nil
}

func (m *MockTape) Set(pos int, symbols ...turing.Symbol) error {
	m.setPos = pos
	m.setSymbols = symbols
	return nil
}

func TestHead(t *testing.T) {
	t.Run("Attach", func(t *testing.T) {
		t.Log("should attach to a tape in a specified position")

		tape := MockTape{}

		head := turing.Head{}
		head.Attach(&tape, 10)
		_, err := head.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, 10, tape.getPos)
		}
	})

	t.Run("Read", func(t *testing.T) {
		t.Log("shoulw read what is on the tape on the current position")
		tape := turing.NewInfiniteTape()

		tape.Set(5, "bananas")

		head := turing.Head{}

		head.Attach(tape, 5)

		v, err := head.Read()

		if assert.NoError(t, err) {
			assert.Equal(t, "bananas", v)
		}

	})

	t.Run("Write", func(t *testing.T) {
		t.Log("shoulw set value on the tape for the current position")
		tape := &MockTape{}

		tape.Set(5, "bananas")

		head := turing.Head{}

		head.Attach(tape, 5)

		err := head.Write("potatoes")
		if assert.NoError(t, err) {
			assert.Equal(t, 5, tape.setPos)
			assert.Equal(t, []turing.Symbol{"potatoes"}, tape.setSymbols)
		}
	})

	t.Run("Left", func(t *testing.T) {
		t.Log("should move left on the tape")

		tape := MockTape{}

		head := turing.Head{}
		head.Attach(&tape, 10)
		head.Move(turing.LEFT)
		_, err := head.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, 9, tape.getPos)
		}
	})

	t.Run("Right", func(t *testing.T) {
		t.Log("should move right on the tape")

		tape := MockTape{}

		head := turing.Head{}
		head.Attach(&tape, 10)
		head.Move(turing.RIGHT)
		_, err := head.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, 11, tape.getPos)
		}
	})

	t.Run("Stay", func(t *testing.T) {
		t.Log("should stay on the same place on the tape")

		tape := MockTape{}

		head := turing.Head{}
		head.Attach(&tape, 10)
		head.Move(turing.STAY)
		_, err := head.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, 10, tape.getPos)
		}

		head.Move("bananas")
		_, err = head.Read()
		if assert.NoError(t, err) {
			assert.Equal(t, 10, tape.getPos)
		}
	})
}
