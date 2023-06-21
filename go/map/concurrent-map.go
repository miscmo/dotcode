package main

import (
	//"github.com/orcaman/concurrent-map/v2"
	"sync"
)

func main() {
	// Create a new map.
	//m := cmap.New[string]()
	//
	//// Sets item within map, sets "bar" under key "foo"
	//m.Set("foo", "bar")
	//
	//// Retrieve item from map.
	//bar, ok := m.Get("foo")
	//
	//// Removes item under key "foo"
	//m.Remove("foo")

	sMap := sync.Map{}

}
