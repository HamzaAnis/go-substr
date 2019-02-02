package substr

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		in       string
		number   int
		expected string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 0, "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 10, "KLMNOPQRSTUVWXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 4, "EFGHIJKLMNOPQRSTUVWXYZ012345"},
		{"OK", 3, ""},
		{"OK", 1, "K"},
		{"OK", 2, ""},	
	}

	for i, test := range tests {
		var data = String{test.in}
		actual := data.Substr(test.number)
		if actual != test.expected {
			t.Errorf("#%d: Substr(%v)=%s; expected %s", i, test.number, actual, test.expected)
		}
	}
}

func TestSubstrMulti(t *testing.T) {
	tests := []struct {
		in       string
		number1  int
		number2  int
		expected string
	}{
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 0, 6, "ABCDEF"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 10, 6, "KLMNOP"},
	}

	for i, test := range tests {
		var data = String{test.in}
		actual := data.Substr(test.number1, test.number2)
		if actual != test.expected {
			t.Errorf("#%d: Substr(%v,%v)=%s; expected %s", i, test.number1, test.number2, actual, test.expected)
		}
	}
}
