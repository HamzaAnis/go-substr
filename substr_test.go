package substr

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"OK", "OK"},
		{"OK", "OK"},
		{"OK", "OK"},
		{"OK", "OK"},
		{"OK", "OK"},
	}

	for i, test := range tests {
		actual := Substr(1)
		if actual != test.expected {
			t.Errorf("#%d: Substr(%v)=%s; expected %s", i, test.in, actual, test.expected)
		}
	}
}
