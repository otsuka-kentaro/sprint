package sprint

import (
	"fmt"
	"testing"
)

func TestSprint(t *testing.T) {
	type testType int
	var ni interface{}
	zero := 0
	a := "a"
	bl := true
	strct := struct {
		in  int
		pin *int
		st  string
		pst *string
		bl  bool
		pbl *bool
		ni  interface{}
		pni *interface{}
		skipped interface{}
		pskipped *interface{}
		testTyp testType
		pTestTyp *testType
	}{
		in:  zero,
		pin: &zero,
		st:  a,
		pst: &a,
		bl:  bl,
		pbl: &bl,
		ni:  ni,
		pni: &ni,
	}
	// contains no newline characters
	strctExpect := fmt.Sprintf("{"+
		"in: %d, "+
		"pin: %d, "+
		"st: %s, "+
		"pst: %s, "+
		"bl: %t, "+
		"pbl: %t, "+
		"testTyp: 0" +
		"}", zero, zero, a, a, bl, bl)

	tests := []struct {
		name   string      // test case name
		input  interface{} // test target object
		expect string      // expect string
	}{
		// data
		{"nil", ni, "<nil>"},
		{"int", zero, "0"},
		{"string", a, "a"},
		{"bool", bl, "true"},
		{"struct", strct, strctExpect},

		// pointer
		{"*nil", &ni, "<nil>"},
		{"*int", &zero, "0"},
		{"*string", &a, "a"},
		{"*bool", &bl, "true"},
		{"*struct", &strct, strctExpect},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
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
