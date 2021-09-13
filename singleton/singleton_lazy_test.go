package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance3(t *testing.T) {
	i3 := GetInstance3()
	fmt.Printf("%#v\n", i3)
}

func TestGetInstance4(t *testing.T) {
	i4 := GetInstance4()
	fmt.Printf("%#v\n", i4)
}

func TestGetInstance5(t *testing.T) {
	i5 := GetInstance5()
	fmt.Printf("%#v\n", i5)
}

func TestGetInstance6(t *testing.T) {
	i6 := GetInstance6()
	fmt.Printf("%#v\n", i6)
}

func TestGetInstance7(t *testing.T) {
	i7 := GetInstance7()
	fmt.Printf("%#v\n", i7)
}
