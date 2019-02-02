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
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 40, ""},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -10, "WXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -11, "VWXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -50, "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"},
		{"OK", 3, ""},
		{"OK", 1, "K"},
		{"OK", 2, ""},
	}

	for i, test := range tests {
		var data = String{test.in}
		actual := data.Substr(test.number)
		if actual != test.expected {
			t.Errorf("#%d: %v.Substr(%v)=%s; expected %s", i, test.in, test.number, actual, test.expected)
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
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 5, 10, "FGHIJKLMNO"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 40, 6, ""},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -10, 5, "WXYZ0"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -10, 10, "WXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -10, 11, "WXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -10, 20, "WXYZ012345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -5, 5, "12345"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", -30, 20, "CDEFGHIJKLMNOPQRSTUV"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 5, -6, "FGHIJKLMNOPQRSTUVWXYZ01234"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 5, -20, "FGHIJKLMNOPQ"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ012345", 5, -32, ""},
	}

	for i, test := range tests {
		var data = String{test.in}
		actual := data.Substr(test.number1, test.number2)
		if actual != test.expected {
			t.Errorf("#%d: %v.Substr(%v,%v)=%s; expected %s", i, test.in, test.number1, test.number2, actual, test.expected)
		}
	}
}
