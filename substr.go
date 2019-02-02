package substr

import "fmt"

// This struct holds the value.
type String struct {
	data string
}

// Substr clone of https://help.adobe.com/en_US/FlashPlatform/reference/actionscript/3/String.html#substr()
func (v *String) Substr(n int) string {
	fmt.Println(v.data)
	return "OK"
}
