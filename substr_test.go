package substr

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		in       string
		number 	int
		expected string
	}{
		{"OK",0, "OK"},
		{"OK", 10,"OK"},
		{"OK", -10,"OK"},
		}

	for i, test := range tests {
		var data = String{"Ok"}
		actual := data.Substr(1)
		if actual != test.expected {
			t.Errorf("#%d: Substr(%v)=%s; expected %s", i, test.in, actual, test.expected)
		}
	}
}
