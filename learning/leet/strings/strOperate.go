package strings

import (
	"bytes"
	"fmt"
)

func TestString() {
	fmt.Println(replaceSpace("We are happy."))
}

func replaceSpace(strOri string) string {
	var buff bytes.Buffer
	for _, ch := range strOri {
		if ch == ' ' {
			buff.WriteString("%20")
		} else{
			buff.WriteRune(ch)
		}
	}

	return buff.String()
}
