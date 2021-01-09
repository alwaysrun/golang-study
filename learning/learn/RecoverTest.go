package learn

import (
	"errors"
	"fmt"
	"runtime"
)

func RecoverTest() (err error)  {
	defer func() {
		if r:=recover(); r!=nil{
			switch x:=r.(type) {
			case runtime.Error:
				err = x
			case error:
				err = x
			case string:
				err = errors.New(x)
			default:
				err = fmt.Errorf("Unknown panic %v", x)
			}
		}
	}()

	panic("Test panic")
}
