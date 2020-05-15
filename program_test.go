package turing_test

import (
	"fmt"
	"testing"

	"github.com/massahud/turing"
	"github.com/stretchr/testify/assert"
)

func TestProgram(t *testing.T) {
	t.Run("AddAndList", func(t *testing.T) {
		t.Log("should be able to add and list operations")

		state := turing.State{"state", false}
		halt := turing.State{"halt", true}

		op1 := turing.Op{state, 0, 0, turing.RIGHT, state}
		op2 := turing.Op{state, 1, 0, turing.RIGHT, state}
		op3 := turing.Op{state, nil, nil, turing.STAY, halt}

		program := turing.Program{}
		program.AddOp(op1)
		program.AddOp(op2)
		program.AddOp(op3)

		ops := program.ListOps()

		assert.Len(t, ops, 3)
		assert.Contains(t, ops, op1)
		assert.Contains(t, ops, op2)
		assert.Contains(t, ops, op3)

	})

	t.Run("FindOp", func(t *testing.T) {
		t.Log("should be able to find operations")

		state := turing.State{"state", false}
		halt := turing.State{"halt", true}

		op1 := turing.Op{state, 0, 0, turing.RIGHT, state}
		op2 := turing.Op{state, 1, 0, turing.RIGHT, state}
		op3 := turing.Op{state, nil, nil, turing.STAY, halt}

		program := turing.Program{}
		program.AddOp(op1)
		program.AddOp(op2)
		program.AddOp(op3)

		op, err := program.FindOp(state, 0)
		if assert.NoError(t, err) {
			assert.Equal(t, op1, op)
		}

		op, err = program.FindOp(state, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, op2, op)
		}

		op, err = program.FindOp(state, nil)

		if assert.NoError(t, err) {
			assert.Equal(t, op3, op)
		}
	})

	t.Run("NoState", func(t *testing.T) {
		t.Log("should return error when there is no op for the state")

		state := turing.State{"state", false}
		dummy := turing.State{"dummy", false}

		program := turing.Program{}
		program.AddOp(turing.Op{state, 0, 0, turing.RIGHT, state})

		_, err := program.FindOp(dummy, 0)

		assert.EqualError(t, err, fmt.Sprintf("no operation for state %v", dummy))

	})

	t.Run("NoSymbol", func(t *testing.T) {
		t.Log("should return error when there is no op for the symbol")

		state := turing.State{"state", false}

		program := turing.Program{}
		program.AddOp(turing.Op{state, 0, 0, turing.RIGHT, state})

		_, err := program.FindOp(state, 1)

		assert.EqualError(t, err, fmt.Sprintf("no operation for state %v and symbol %v", state, 1))

	})

	t.Run("AnySymbol", func(t *testing.T) {
		t.Log("should return operation when there is a match for any symbol")

		state := turing.State{"state", false}

		op1 := turing.Op{state, turing.ANY, 0, turing.RIGHT, state}

		program := turing.Program{}
		program.AddOp(op1)

		op, err := program.FindOp(state, nil)
		if assert.NoError(t, err) {
			assert.Equal(t, op, op1)
		}

		op, err = program.FindOp(state, 1)
		if assert.NoError(t, err) {
			assert.Equal(t, op, op1)
		}

		op, err = program.FindOp(state, 2)
		if assert.NoError(t, err) {
			assert.Equal(t, op, op1)
		}
	})

}
