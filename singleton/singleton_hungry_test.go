package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance1(t *testing.T) {
	i := GetInstance1()
	fmt.Printf("%#v\n", i)
}

func TestGetInstance2(t *testing.T) {
	i := GetInstance2()
	fmt.Printf("%#v\n", i)
}
