// Package turing implements a turing machine.
//
// A turing machine is composed of a Head attached to a Tape
// and a Program, which is a list of Operation's. It also tracks
// the current state.
//
// To run a program, you attach the head to a pre initialized tape and
// executes the program given the initial state of the machine.
// Each position of the tape can have a symbol or be blank (nil).
//
// The machine tries to find the operation that should be executed
// for the current state and current symbol under the head.
//
// Each operation can change the current symbol, then move the head left or right,
// and change the current machine state after the head moves. It repeats this
// until it finds a halting state.
//
// To simplify the turing program, operations can have some special definitions,
// it can match ANY symbol or KEEP the current symbol. The head also can STAY
// on the same position.
package turing
