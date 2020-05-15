package main

import (
	"fmt"
	"syscall/js"

	"github.com/massahud/turing"
)

// createSeparate01 creates a simple program that separates 0's from 1's
// of the given sequence.
//
// It returns the start state and the program.
func createSeparate01() (turing.State, *turing.Program) {
	get1 := turing.State{Name: "get1"}
	get0 := turing.State{Name: "get0"}
	back0 := turing.State{Name: "back0"}
	back1 := turing.State{Name: "back1"}
	halt := turing.State{Name: "halt", Halt: true}

	program := turing.Program{}
	program.AddOp(turing.Op{State: get1, Symbol: 1, WriteSymbol: nil, Movement: turing.RIGHT, NextState: get0})
	program.AddOp(turing.Op{State: get1, Symbol: 0, WriteSymbol: 0, Movement: turing.RIGHT, NextState: get1})
	program.AddOp(turing.Op{State: get1, Symbol: nil, WriteSymbol: nil, Movement: turing.STAY, NextState: halt})
	program.AddOp(turing.Op{State: get0, Symbol: 1, WriteSymbol: 1, Movement: turing.RIGHT, NextState: get0})
	program.AddOp(turing.Op{State: get0, Symbol: 0, WriteSymbol: 1, Movement: turing.LEFT, NextState: back0})
	program.AddOp(turing.Op{State: get0, Symbol: nil, WriteSymbol: nil, Movement: turing.LEFT, NextState: back1})
	program.AddOp(turing.Op{State: back0, Symbol: turing.ANY, WriteSymbol: turing.KEEP, Movement: turing.LEFT, NextState: back0})
	program.AddOp(turing.Op{State: back0, Symbol: nil, WriteSymbol: 0, Movement: turing.RIGHT, NextState: get1})
	program.AddOp(turing.Op{State: back1, Symbol: turing.ANY, WriteSymbol: turing.KEEP, Movement: turing.LEFT, NextState: back1})
	program.AddOp(turing.Op{State: back1, Symbol: nil, WriteSymbol: 1, Movement: turing.STAY, NextState: halt})

	return get1, &program
}

func main() {

	head := turing.Head{}
	start, program := createSeparate01()
	machine := turing.Machine{Head: &head, Program: program, State: start}

	jsSequence := js.Global().Get("sequence")

	if !jsSequence.InstanceOf(js.Global().Get("Array")) {
		fmt.Printf("ERROR: sequence is not an Array: %v\n", jsSequence)
		return
	}

	sequence := make([]turing.Symbol, jsSequence.Length())
	for i := range sequence {
		sequence[i] = jsSequence.Index(i).Int()
	}

	fmt.Println("JS global 'sequence': ", sequence)

	tape := turing.NewInfiniteTape()
	tape.Set(0, sequence...)

	head.Attach(tape, 0)

	fmt.Println("Before:", len(sequence))
	fmt.Println(head.PrintTape(0, len(sequence)))

	err := machine.Run()
	if err != nil {
		fmt.Println("ERROR: error executing machine:", err.Error())
		return
	}

	fmt.Println("After:")
	fmt.Println(head.PrintTape(0, len(sequence)))

}
