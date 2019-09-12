package sprint

import (
	"fmt"
	"testing"
)

func TestSprint(t *testing.T) {
	zero := 0
	a := "a"
	bl := true
	strct := struct{
		in int
		pin *int
		st string
		pst *string
		bl bool
		pbl *bool
	}{
		in: zero,
		pin: &zero,
		st: a,
		pst: &a,
		bl: bl,
		pbl: &bl,
	}
	// contains no newline characters
	strctExpect := fmt.Sprintf("{" +
		"in: %d, " +
		"pin: %d, " +
		"st: %s, " +
		"pst: %s, " +
		"bl: %t, " +
		"pbl: %t" +
		"}", zero, zero, a, a, bl, bl)

	tests := []struct {
		name string	// test case name
		input interface{} // test target object
		expect string // expect string
	}{
		// data
		{"int", zero, "0"},
		{"string", a, "a"},
		{"bool", bl, "true"},
		{"struct", strct, strctExpect},

		// pointer
		{"*int", &zero, "0"},
		{"*string", &a, "a"},
		{"*bool", &bl, "true"},
		{"struct", &strct, strctExpect},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T){
			defer func() {
				if err := recover(); err != nil {
					tt.Error(err)
				}
			}()
			s := Sprint(test.input)
			if s != test.expect {
				tt.Errorf("expected: %s. but got actual: %s", test.expect, s)
			}
		})
	}
}
