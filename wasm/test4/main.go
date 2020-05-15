package main

import (
	"fmt"
	"strings"
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

	exit := make(chan struct{})

	js.Global().Set("runMachine", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		jsSequence := args[0]
		defer func() {
			exit <- struct{}{}
		}()

		if !jsSequence.InstanceOf(js.Global().Get("Array")) {
			return fmt.Sprintf("ERROR: sequence is not an Array: %v\n", jsSequence)
		}

		sequence := make([]turing.Symbol, jsSequence.Length())
		for i := range sequence {
			sequence[i] = jsSequence.Index(i).Int()
		}

		builder := strings.Builder{}
		builder.WriteString(fmt.Sprintf("Running for sequence: %v\n\n", sequence))

		tape := turing.NewInfiniteTape()
		tape.Set(0, sequence...)

		head.Attach(tape, 0)

		builder.WriteString(fmt.Sprintln("Initial configuration:", len(sequence)))
		builder.WriteString(head.PrintTape(0, len(sequence)))

		err := machine.Run()
		if err != nil {
			return fmt.Sprint("ERROR: error executing machine:", err.Error())
		}

		builder.WriteString("After execution:\n")
		builder.WriteString(head.PrintTape(0, len(sequence)))
		return builder.String()
	}))

	<-exit
}
