package substr

// This struct holds the value.
type String struct {
	data string
}

// Substr clone of https://help.adobe.com/en_US/FlashPlatform/reference/actionscript/3/String.html#substr()
func (v *String) Substr(n1 int, n2 ...int) string {
	substr := ""
	if len(n2) <= 0 {
		if n1 >= 0 && n1 <= len(v.data) {
			substr = v.data[n1:]
		}
	} else {
		if n1 >= 0 && n1+n2[0] <= len(v.data) {
			substr = v.data[n1 : n1+n2[0]]
		}
	}
	return substr
}
