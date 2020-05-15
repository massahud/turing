package turing_test

import (
	"fmt"

	"github.com/massahud/turing"
)

// Simple machine that transforms everything in zeros and halt
// on blank (nil).
func Example_zeroAll() {
	// states
	zeroing := turing.State{"zero all", false}
	halt := turing.State{"halt", true}

	zeroAll := turing.Op{zeroing, turing.ANY, 0, turing.RIGHT, zeroing}
	finishOnBlank := turing.Op{zeroing, nil, nil, turing.STAY, halt}

	program := turing.Program{}
	program.AddOp(zeroAll)
	program.AddOp(finishOnBlank)

	head := turing.Head{}

	machine := turing.Machine{Head: &head, Program: &program, State: zeroing}

	tape := turing.NewInfiniteTape()
	tape.Set(0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1)

	head.Attach(tape, 0)

	machine.Run()

	fmt.Println("Head:", head.Pos())
	fmt.Println("State:", machine.State)
	fmt.Printf("Tape[%d,%d]:", head.MinPos(), head.MaxPos())
	for i := head.MinPos(); i <= head.MaxPos(); i++ {
		v, _ := tape.Get(i)
		fmt.Printf(" %v", v)
	}

	// Output:
	// Head: 10
	// State: [halt]
	// Tape[0,10]: 0 0 0 0 0 0 0 0 0 0 <nil>
}

// Mirrors a [01]* string.
func Example_mirror() {
	// states
	// goes to end
	toEnd := turing.State{"toEnd0", false}
	// cut current symbol of source sequence
	cut := turing.State{"copy", false}
	// pastes one of the symbols at the end
	paste0 := turing.State{"paste0", false}
	paste1 := turing.State{"paste1", false}
	// goes to where it cut the sequence sequence to fix it
	return0 := turing.State{"return0", false}
	return1 := turing.State{"return1", false}
	// halt
	halt := turing.State{"halt", true}

	program := turing.Program{}

	program.AddOp(turing.Op{toEnd, nil, turing.KEEP, turing.LEFT, cut})
	program.AddOp(turing.Op{toEnd, turing.ANY, turing.KEEP, turing.RIGHT, toEnd})

	program.AddOp(turing.Op{cut, nil, turing.KEEP, turing.RIGHT, halt})
	program.AddOp(turing.Op{cut, 0, nil, turing.RIGHT, paste0})
	program.AddOp(turing.Op{cut, 1, nil, turing.RIGHT, paste1})

	program.AddOp(turing.Op{paste0, nil, 0, turing.STAY, return0})
	program.AddOp(turing.Op{paste0, turing.ANY, turing.KEEP, turing.RIGHT, paste0})

	program.AddOp(turing.Op{paste1, nil, 1, turing.STAY, return1})
	program.AddOp(turing.Op{paste1, turing.ANY, turing.KEEP, turing.RIGHT, paste1})

	program.AddOp(turing.Op{return0, nil, 0, turing.LEFT, cut})
	program.AddOp(turing.Op{return0, turing.ANY, turing.KEEP, turing.LEFT, return0})

	program.AddOp(turing.Op{return1, nil, 1, turing.LEFT, cut})
	program.AddOp(turing.Op{return1, turing.ANY, turing.KEEP, turing.LEFT, return1})

	head := turing.Head{}

	machine := turing.Machine{Head: &head, Program: &program, State: toEnd}

	tape := turing.NewInfiniteTape()
	tape.Set(0, 1, 0, 1)

	head.Attach(tape, 0)

	err := machine.Run()
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	fmt.Println("Head:", head.Pos())
	fmt.Println("State:", machine.State)
	fmt.Printf("Tape[%d,%d]:", head.MinPos(), head.MaxPos())

	for i := head.MinPos(); i <= head.MaxPos(); i++ {
		v, _ := tape.Get(i)
		fmt.Printf(" %v", v)
	}

	// Output:
	// Head: 0
	// State: [halt]
	// Tape[-1,5]: <nil> 1 0 1 1 0 1
}

// Copies a [01]* string, leaving a blank between the original and the copy.
// Head stops in the middle blank position
func Example_copy() {

	// states
	// cut current symbol of source sequence
	cut := turing.State{"copy", false}
	// goes to end of source with one of the symbols
	toEnd0 := turing.State{"toEnd0", false}
	toEnd1 := turing.State{"toEnd1", false}
	// write one of the symbols at the end of destination sequence
	write0 := turing.State{"write0", false}
	write1 := turing.State{"write1", false}
	// go to start of destination sequence after writing
	wrote0 := turing.State{"wrote0", false}
	wrote1 := turing.State{"wrote1", false}
	// go to where it cut source sequence to fix it
	return0 := turing.State{"return0", false}
	return1 := turing.State{"return1", false}
	// halt
	halt := turing.State{"halt", true}

	program := turing.Program{}
	program.AddOp(turing.Op{cut, nil, nil, turing.STAY, halt})
	program.AddOp(turing.Op{cut, 0, nil, turing.RIGHT, toEnd0})
	program.AddOp(turing.Op{cut, 1, nil, turing.RIGHT, toEnd1})

	program.AddOp(turing.Op{toEnd0, nil, nil, turing.RIGHT, write0})
	program.AddOp(turing.Op{toEnd0, turing.ANY, turing.KEEP, turing.RIGHT, toEnd0})

	program.AddOp(turing.Op{toEnd1, nil, nil, turing.RIGHT, write1})
	program.AddOp(turing.Op{toEnd1, turing.ANY, turing.KEEP, turing.RIGHT, toEnd1})

	program.AddOp(turing.Op{write0, nil, 0, turing.LEFT, wrote0})
	program.AddOp(turing.Op{write0, turing.ANY, turing.KEEP, turing.RIGHT, write0})

	program.AddOp(turing.Op{write1, nil, 1, turing.LEFT, wrote1})
	program.AddOp(turing.Op{write1, turing.ANY, turing.KEEP, turing.RIGHT, write1})

	program.AddOp(turing.Op{wrote0, nil, nil, turing.LEFT, return0})
	program.AddOp(turing.Op{wrote0, turing.ANY, turing.KEEP, turing.LEFT, wrote0})

	program.AddOp(turing.Op{wrote1, nil, nil, turing.LEFT, return1})
	program.AddOp(turing.Op{wrote1, turing.ANY, turing.KEEP, turing.LEFT, wrote1})

	program.AddOp(turing.Op{return0, nil, 0, turing.RIGHT, cut})
	program.AddOp(turing.Op{return0, turing.ANY, turing.KEEP, turing.LEFT, return0})

	program.AddOp(turing.Op{return1, nil, 1, turing.RIGHT, cut})
	program.AddOp(turing.Op{return1, turing.ANY, turing.KEEP, turing.LEFT, return1})

	head := turing.Head{}

	machine := turing.Machine{Head: &head, Program: &program, State: cut}

	tape := turing.NewInfiniteTape()
	tape.Set(0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 1)

	head.Attach(tape, 0)

	err := machine.Run()
	if err != nil {
		fmt.Println("error:", err.Error())
	}

	fmt.Println("Head:", head.Pos())
	fmt.Println("State:", machine.State)
	fmt.Printf("Tape[%d,%d]:", head.MinPos(), head.MaxPos())
	for i := head.MinPos(); i <= head.MaxPos(); i++ {
		v, _ := tape.Get(i)
		fmt.Printf(" %v", v)
	}

	// Output:
	// Head: 10
	// State: [halt]
	// Tape[0,20]: 1 1 1 0 1 1 0 1 0 1 <nil> 1 1 1 0 1 1 0 1 0 1
}
