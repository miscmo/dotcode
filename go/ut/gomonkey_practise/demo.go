package gomonkey_practise

type Fake struct {
}

func (f *Fake) Ok() bool {
	return true
}

func (f *Fake) Get() string {
	return "hh"
}

/*
 *
 */

func UseFake1() (bool, string) {

	fake1 := Fake{}

	return fake1.Ok(), fake1.Get()
}

func UseFake2()(bool, string) {
	fake2 := Fake{}

	return fake2.Ok(), fake2.Get()
}