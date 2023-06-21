package gomonkey_practise

import (
	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockerFunc func() func()

var FakeOkReturnFalse MockerFunc = func() func() {
	patch := gomonkey.ApplyMethod(&Fake{}, "Ok", func() bool {
		return false
	})

	return func() {
		patch.Reset()
	}
}

var FakeGetReturnXX MockerFunc = func() func() {
	patch := gomonkey.ApplyMethod(&Fake{}, "Get", func() string {
		return "XX"
	})

	return func() {
		patch.Reset()
	}
}

func TestUseFake1(t *testing.T) {

	cases := []struct{
		Name string
		ExpectedIsOk bool
		ExpectedStr string
		Mockers []MockerFunc
	} {
		{"TestDefaultValue", true, "hh", nil},
		{"TestOkFalseStrXX", false, "XX", []MockerFunc{FakeOkReturnFalse, FakeGetReturnXX}},
	}

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			for _, mocker := range v.Mockers {
				defer mocker()()
			}

			ok, str := UseFake1()

			assert.Equal(t, v.ExpectedIsOk, ok)
			assert.Equal(t, v.ExpectedStr, str)
		})
	}
}
func TestUseFake2(t *testing.T) {

	cases := []struct{
		Name string
		ExpectedIsOk bool
		ExpectedStr string
		Mockers []MockerFunc
	} {
		{"TestDefaultValue", true, "hh", nil},
		{"TestOkFalseStrXX", false, "XX", []MockerFunc{FakeOkReturnFalse, FakeGetReturnXX}},
	}

	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			for _, mocker := range v.Mockers {
				defer mocker()()
			}

			ok, str := UseFake2()

			assert.Equal(t, v.ExpectedIsOk, ok)
			assert.Equal(t, v.ExpectedStr, str)
		})
	}
}
