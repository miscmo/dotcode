package main

import "github.com/miscmo/dotcode/go/design_pattern/abstract_factory/pipeline"

func main() {
	p := pipeline.Of(pipeline.DefaultConfig())
	p.Exec()
}
