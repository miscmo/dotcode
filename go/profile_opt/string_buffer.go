package main

import (
	"bytes"
	"fmt"
)

func Process(data []string) string {
	var buf bytes.Buffer

	for _, item := range data {
		buf.WriteString(item)

		buf.WriteString("\n")
	}

	return buf.String()
}

func NewString() string {
	data := make([]string, 100000)

	for i := 0; i < len(data); i++ {
		data[i] = fmt.Sprintf("Item %d", i)
	}

	result := Process(data)

	return result
}

func main() {
	NewString()
}