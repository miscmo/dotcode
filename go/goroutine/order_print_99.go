package main

import "strconv"

type pp struct {
	i string
	v int
	b *string
	d int
}

func main() {
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(i)

		p := pp {
			i: str,
			v: i,
			b: &str,
			d: i,
		}

		go func() {
			MyPrint(&p)
		}()
	}
}


func MyPrint(p *pp) {
	println((p.d))
}