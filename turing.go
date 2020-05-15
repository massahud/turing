package turing

import "fmt"

const (
	// LEFT movement
	LEFT = "left"
	// RIGHT movement
	RIGHT = "right"
	// STAY movement
	STAY = "stay"

	// ANY symbol
	ANY = "__turing[any]"
	// KEEP symbol
	KEEP = "__turing[keep]"
)

// State identifies a machine state
// it can also be a halt state, which halts the machine
type State struct {
	// Name is the state name
	Name string
	// Halt informs if the state is a halting state
	Halt bool
}

// String writes the state name. If it is a halt state
// it surrounds the name with brackets.
func (s State) String() string {
	if s.Halt {
		return fmt.Sprintf("[%s]", s.Name)
	}
	return s.Name
}

// Symbol is a turing machine symbol. It can be anything.
//
// nil is the blank symbol
type Symbol interface{}

// Tape is a turing machine tape.
type Tape interface {
	Get(pos int) (Symbol, error)
	Set(pos int, symbols ...Symbol) error
}

// Machine is a turing machine, it has a head, a program to execute and the
// initial state.
type Machine struct {
	Head    *Head
	Program *Program
	State   State
}

// Step executes one step of the machine
func (m *Machine) Step() error {
	if m.State.Halt {
		return fmt.Errorf("machine is halted, state %s", m.State.String())
	}
	v, err := m.Head.Read()
	oper, err := m.Program.FindOp(m.State, v)
	if err != nil {
		return err
	}

	if oper.WriteSymbol != KEEP {
		m.Head.Write(oper.WriteSymbol)
	}
	m.Head.Move(oper.Movement)
	m.State = oper.NextState
	return nil
}

// Run executes the current program until it reaches a halt state or there is no
// operation for current state and symbol under head.
func (m *Machine) Run() error {

	for instr := 1; !m.State.Halt; instr++ {
		err := m.Step()
		if err != nil {
			return fmt.Errorf("Error at instruction %d: %s", instr, err.Error())
		}
	}
	return nil
}
