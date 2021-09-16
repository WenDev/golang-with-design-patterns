package factory

import (
	"fmt"
	"testing"
)

func TestHumanFactory_CreateHuman(t *testing.T) {
	// 声明阴阳八卦炉
	yinYangLu := &HumanFactory{}
	// 女娲第一次造人,火候不足,于是白人产生了
	fmt.Println("--造出的第一批人是白色人种--")
	wh := yinYangLu.CreateHuman("w")
	wh.GetColor()
	wh.Talk()
	// 女娲第二次造人,火候过足,于是黑人产生了
	fmt.Println("--造出的第二批人是黑色人种--")
	bh := yinYangLu.CreateHuman("b")
	bh.GetColor()
	bh.Talk()
	// 第三次造人,火候刚刚好,于是黄色人种产生了
	fmt.Println("--造出的第三批人是黄色人种--")
	yh := yinYangLu.CreateHuman("y")
	yh.GetColor()
	yh.Talk()
}
