package main

import "qiniupkg.com/x/errors.v7"

type UndoIntSet struct {
	IntSet
	functions []func()
}

func NewUndoIntSet() UndoIntSet {
	return UndoIntSet{NewIntSet(), nil}
}

// Override
func (set *UndoIntSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() { set.Delete(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

// Override
func (set *UndoIntSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.functions = append(set.functions, func() { set.Add(x) })
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil
	}
	set.functions = set.functions[:index]
	return nil
}
