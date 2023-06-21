package main

import "errors"

type IntSet struct {
	data map[int]bool
}

func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(x int) {
	set.data[x] = true
}

func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}

type UndoableIntSet struct {
	IntSet
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

func (set *UndoableIntSet) Add(x int) {
	if !set.Contains(x) {
		set.IntSet.Add(x)
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) {
	set.IntSet.Delete(x)
	if set.Contains(x) {
		set.IntSet.Delete(x)
		set.functions = append(set.functions, func() {
			set.Add(x)
		})
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}

	index := len(set.functions) - 1

	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil	// for gc
	}

	set.functions = set.functions[:index]
	return nil
}


type Undo []func()

