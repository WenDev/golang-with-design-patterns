package factory

import (
	"fmt"
	"testing"
)

func TestCreateHuman1(t *testing.T) {
	// 女娲第一次造人,火候不足,于是白人产生了
	fmt.Println("--造出的第一批人是白色人种--")
	whf := WhiteHumanFactory{}
	wh := whf.CreateHuman()
	wh.GetColor()
	wh.Talk()
	// 女娲第二次造人,火候过足,于是黑人产生了
	fmt.Println("--造出的第二批人是黑色人种--")
	bhf := BlackHumanFactory{}
	bh := bhf.CreateHuman()
	bh.GetColor()
	bh.Talk()
	// 第三次造人,火候刚刚好,于是黄色人种产生了
	fmt.Println("--造出的第三批人是黄色人种--")
	yhf := YellowHumanFactory{}
	yh := yhf.CreateHuman()
	yh.GetColor()
	yh.Talk()
}
