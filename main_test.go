// main_test.go
package main

import "testing"

func TestGoatLatin(t *testing.T) {
	input := "Words written in memory"
	output := "ordsWmaa rittenwmaaa inmaaaa emorymmaaaaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}
	input = "i"
	output = "imaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}

	input = "la"
	output = "almaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}

	input = "e"
	output = "emaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}
	input = "run in a circle"
	output = "unrmaa inmaaa amaaaa irclecmaaaaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}
	input = ""
	output = ""
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}
	input = "this  has  many   spaces"
	output = "histmaa ashmaaa anymmaaaa pacessmaaaaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}

	input = "How? Is this the first time?"
	output = "ow?Hmaa Ismaaa histmaaaa hetmaaaaa irstfmaaaaaa ime?tmaaaaaaa"
	if goatLatin(input) != output {
		t.Log("expecting:", output)
		t.Fail()
	}
	//         t.Errorf("Abs(-1) = %d; want 1", got)
}

func BenchmarkGoatLatin(b *testing.B) {
	s := "Words written in memory"
	for i := 0; i<= 10000; i++ {
			goatLatin(s)
	}

	//         t.Errorf("Abs(-1) = %d; want 1", got)
}
