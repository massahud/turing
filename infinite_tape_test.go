package turing_test

import (
	"testing"

	"github.com/massahud/turing"
	"github.com/stretchr/testify/assert"
)

func TestInfiniteTape(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		tape := turing.NewInfiniteTape()
		assert.NotNil(t, tape)
	})

	t.Run("SetAndGet", func(t *testing.T) {
		t.Log("should set and get on any position")

		tape := turing.NewInfiniteTape()

		tests := []struct {
			pos      int
			existing turing.Symbol
			value    turing.Symbol
		}{
			{
				pos:   0,
				value: "zero",
			},
			{
				pos:   10,
				value: "positive",
			},
			{
				pos:   -100,
				value: "negative",
			},
			{
				pos:      -100,
				existing: "negative",
				value:    "replace",
			},
		}
		for _, tt := range tests {
			value, err := tape.Get(tt.pos)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.existing, value)
			}

			assert.NoError(t, tape.Set(tt.pos, tt.value))

			value, err = tape.Get(tt.pos)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.value, value)
			}
		}
	})

	t.Run("SetArray", func(t *testing.T) {
		t.Log("should set an array of Symbols")

		tape := turing.NewInfiniteTape()
		symbols := []turing.Symbol{"a", "b", "c", "d"}
		tape.Set(0, symbols...)
		for i, s := range symbols {
			v, err := tape.Get(i)
			if assert.NoErrorf(t, err, "%d: %v", i, s) {
				assert.Equal(t, s, v)
			}
		}
	})

	t.Run("SetArrayNegative", func(t *testing.T) {
		t.Log("should set an array of Symbols on negative position")
		tape := turing.NewInfiniteTape()
		symbols := []turing.Symbol{"a", "b", "c", "d"}
		tape.Set(-125, symbols...)
		for i, s := range symbols {
			v, err := tape.Get(i - 125)
			if assert.NoErrorf(t, err, "%d: %v", i-125, s) {
				assert.Equal(t, s, v)
			}
		}
	})

	t.Run("SetOverwriteArray", func(t *testing.T) {
		t.Log("should overwrite symbols with array")
		tape := turing.NewInfiniteTape()
		tape.Set(0, "a", "b", "c", "d", "e")
		tape.Set(-3, "1", "2", "3", "4", "5")

		for i, s := range []turing.Symbol{"1", "2", "3", "4", "5", "c", "d", "e"} {
			v, err := tape.Get(i - 3)
			if assert.NoErrorf(t, err, "%d: %v", i-3, s) {
				assert.Equal(t, s, v)
			}
		}
	})

}
