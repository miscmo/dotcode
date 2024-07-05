package main

import (
	"fmt"
	"sync"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type TNoCopy struct {
	// noCopy noCopy
	T    string
	tmux sync.Mutex
}

func main() {
	t := TNoCopy{T: "test"}
	t.tmux.Lock()

	ct := t
	ct.tmux.Lock()
	fmt.Printf("t: %s\n", t.T)
	fmt.Printf("ct: %s\n", ct.T)

	// var mu sync.Mutex

	// var i int

	// mu.Lock()

	// m := mu
	// i += 1
	// m.Unlock()

	// mu.Lock()
	// i += 1
	// mu.Unlock()

}
