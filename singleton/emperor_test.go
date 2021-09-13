package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	for day := 0; day < 3; day++ {
		// 三天见的皇帝都是同一个人,荣幸吧!
		e := GetInstance()
		e.Say()
	}
}
