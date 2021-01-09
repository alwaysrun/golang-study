package number

import "fmt"

func TestNumber()  {
	fmt.Println(isNumber("+100"))
	fmt.Println(isNumber("-1E-16"))
	fmt.Println(isNumber("5.2e2"))
	fmt.Println(isNumber("1a3.14"))
	fmt.Println(isNumber("12e+5.4"))
}

func isNumber(s string) bool {
	//state:=make([]map[byte]int, 0)
	state := []map[byte]int{
		{' ':0, 's':1, 'd':2, '.':4},	 // 0: start with 'blank'
		{'d':2, '.':4},					 // 1: 'sign' before e
		{'d':2, '.':3, 'e':5, ' ':8},	 // 2: 'digit' before '.'
		{'d':3, 'e':5, ' ':8},			 // 3: 'digit' after '.'
		{'d':3},						 // 4: 'digit' after '.'(‘blank’ before 'dot')
		{'s':6, 'd':7},					 // 5: 'e'
		{'d':7},						 // 6: 'sign' after 'e'
		{'d':7, ' ':8},					 // 7: 'digit' after e
		{' ':8},						 // 8: end with 'blank'
	}

	index := 0
	var key byte
	for _,ch := range s{
		if ch>='0' && ch <= '9'{
			key = 'd'
		}else {
			switch ch {
			case '+', '-':
				key = 's'
			case 'e', 'E':
				key = 'e'
			case '.', ' ':
				key = byte(ch)
			default:
				key = '?'
			}
		}

		if _,ok:=state[index][key]; !ok{
			return false
		}
		index = state[index][key]
	}

	switch index {
	case 2,3,7,8:
		return true
	default:
		return false
	}
}
