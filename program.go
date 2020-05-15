package turing

import "fmt"

// Op encapsulates one turing machine operation.
//
// Given a state and the current head symbol, sets the current head symbol,
// move the head and change to another state.
//
// Use ANY on current symbol to make it work for any symbol.
//
// Use KEEP on next symbol to not change current symbol.
type Op struct {
	State       State
	Symbol      Symbol
	WriteSymbol Symbol
	Movement    string
	NextState   State
}

// Program stores the operations, based on current state and symbol under head.
type Program struct {
	ops    map[State]map[Symbol]Op
	length int
}

// FindOp returns the current operation for the State-Symbol tuple.
func (p *Program) FindOp(state State, symbol Symbol) (Op, error) {
	symbolMap := p.ops[state]
	if symbolMap == nil {
		return Op{}, fmt.Errorf("no operation for state %v", state)
	}

	oper, exists := symbolMap[symbol]
	if !exists {
		oper, exists = symbolMap[ANY]
		if !exists {
			return Op{}, fmt.Errorf("no operation for state %v and symbol %v", state, symbol)
		}
	}

	return oper, nil
}

// ListOps returns a list of operations on the machine.
// They are not ordered in any way
func (p *Program) ListOps() []Op {
	opList := make([]Op, 0, p.length)
	for _, symbolMap := range p.ops {
		for _, op := range symbolMap {
			opList = append(opList, op)
		}
	}
	return opList
}

// AddOp adds or rewrite a State-Symbol operation.
func (p *Program) AddOp(op Op) {
	if p.ops == nil {
		p.ops = make(map[State]map[Symbol]Op)
	}
	symbolMap, ok := p.ops[op.State]
	if !ok {
		symbolMap = make(map[Symbol]Op)
		p.ops[op.State] = symbolMap
	}

	if _, ok := symbolMap[op.Symbol]; !ok {
		p.length++
	}
	symbolMap[op.Symbol] = op
}
