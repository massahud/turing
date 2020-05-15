package turing_test

import (
	"testing"

	"github.com/massahud/turing"
	"github.com/stretchr/testify/assert"
)

func TestMachine(t *testing.T) {

	t.Run("StepRight", func(t *testing.T) {
		t.Log("should execute a common step and move right")

		state := turing.State{"state", false}
		newState := turing.State{"newstate", false}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 1)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{state, 1, 0, turing.RIGHT, newState})
		program.AddOp(turing.Op{state, 2, 3, turing.LEFT, newState})

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		if assert.NoError(t, err) {
			v, err := tape.Get(0)
			assert.NoError(t, err)
			assert.Equal(t, 0, v)
			assert.Equal(t, head.Pos(), 1)
			assert.Equal(t, newState, machine.State)
		}
	})

	t.Run("StepLeft", func(t *testing.T) {
		t.Log("should execute a common step and move left")

		state := turing.State{"state", false}
		newState := turing.State{"newstate", false}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 1)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{state, 1, 0, turing.LEFT, newState})
		program.AddOp(turing.Op{state, 2, 3, turing.RIGHT, newState})

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		if assert.NoError(t, err) {
			v, err := tape.Get(0)
			assert.NoError(t, err)
			assert.Equal(t, 0, v)
			assert.Equal(t, head.Pos(), -1)
			assert.Equal(t, newState, machine.State)
		}
	})

	t.Run("StepStay", func(t *testing.T) {
		t.Log("should execute a common step and not move the head")

		state := turing.State{"state", false}
		newState := turing.State{"newstate", false}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 1)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{state, 1, 0, turing.STAY, newState})
		program.AddOp(turing.Op{state, 2, 3, turing.LEFT, newState})

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		if assert.NoError(t, err) {
			v, err := tape.Get(0)
			assert.NoError(t, err)
			assert.Equal(t, 0, v)
			assert.Equal(t, head.Pos(), 0)
			assert.Equal(t, newState, machine.State)
		}
	})

	t.Run("StepKeep", func(t *testing.T) {
		t.Log("should keep the existing value")

		state := turing.State{"state", false}
		newState := turing.State{"newstate", false}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 1)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{state, 1, turing.KEEP, turing.RIGHT, newState})
		program.AddOp(turing.Op{state, 2, 3, turing.LEFT, newState})

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		if assert.NoError(t, err) {
			v, err := tape.Get(0)
			assert.NoError(t, err)
			assert.Equal(t, 1, v)
			assert.Equal(t, head.Pos(), 1)
			assert.Equal(t, newState, machine.State)
		}
	})

	t.Run("StepAny", func(t *testing.T) {
		t.Log("should execute any if no other symbol matches")

		state := turing.State{"state", false}
		newState := turing.State{"newstate", false}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 1)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{state, turing.ANY, 0, turing.RIGHT, newState})
		program.AddOp(turing.Op{state, 2, 3, turing.LEFT, newState})

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		if assert.NoError(t, err) {
			v, err := tape.Get(0)
			assert.NoError(t, err)
			assert.Equal(t, 0, v)
			assert.Equal(t, head.Pos(), 1)
			assert.Equal(t, newState, machine.State)
		}
	})

	t.Run("StepHalted", func(t *testing.T) {
		t.Log("should error when machine is on a halt state")

		halt := turing.State{"banana", true}

		tape := turing.NewInfiniteTape()

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}

		machine := turing.Machine{Head: &head, Program: &program, State: halt}

		err := machine.Step()

		assert.EqualError(t, err, "machine is halted, state [banana]")
	})

	t.Run("StepNoOpState", func(t *testing.T) {
		t.Log("should error when there is no op for current state")
		state := turing.State{"potato", false}

		tape := turing.NewInfiniteTape()

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		assert.EqualError(t, err, "no operation for state potato")
	})

	t.Run("StepNoOpSymbol", func(t *testing.T) {
		t.Log("should error when there is no op for current state-symbol")
		state := turing.State{"potato", false}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 5)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{state, 0, 0, turing.STAY, state})

		machine := turing.Machine{Head: &head, Program: &program, State: state}

		err := machine.Step()

		assert.EqualError(t, err, "no operation for state potato and symbol 5")
	})

	t.Run("Run", func(t *testing.T) {
		t.Log("should run a zero all program")

		zeroAll := turing.State{"zero all", false}
		halt := turing.State{"halt", true}

		tape := turing.NewInfiniteTape()
		tape.Set(0, 1, 0, 1, 0, 1, 1, 1, 1, 0)

		head := turing.Head{}
		head.Attach(tape, 0)

		program := turing.Program{}
		program.AddOp(turing.Op{zeroAll, turing.ANY, 0, turing.RIGHT, zeroAll})
		program.AddOp(turing.Op{zeroAll, nil, nil, turing.STAY, halt})

		machine := turing.Machine{Head: &head, Program: &program, State: zeroAll}

		err := machine.Run()
		if assert.NoError(t, err) {

			assert.Equal(t, 9, head.Pos())

			v, err := head.Read()
			if assert.NoError(t, err) {
				assert.Equal(t, nil, v)
			}

			for i := 0; i < head.Pos(); i++ {
				v, err := tape.Get(i)
				if assert.NoErrorf(t, err, "pos %d", i) {
					assert.Equalf(t, 0, v, "pos %d", i)
				}
			}
		}
	})
}
