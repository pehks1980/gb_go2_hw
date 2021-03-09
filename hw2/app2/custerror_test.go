package custerror

import (
	"fmt"
	"github.com/pehks1980/gb_go2_hw/hw2/app2/custerror"
)

func Example() {
	var s = custerror.New("this is my error with stack")
	fmt.Println(s)
}