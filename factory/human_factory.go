package factory

import "fmt"

// Human 人类总称
type Human interface {
	GetColor() // 每个人种的皮肤都有相应的颜色
	Talk()     // 人类会说话
}

// BlackHuman 黑色人种
type BlackHuman struct{}

func (b *BlackHuman) GetColor() {
	fmt.Println("黑色人种的皮肤颜色是黑色的!")
}

func (b *BlackHuman) Talk() {
	fmt.Println("黑人会说话,一般人听不懂.")
}

// YellowHuman 黄色人种
type YellowHuman struct{}

func (y *YellowHuman) GetColor() {
	fmt.Println("黄色人种的皮肤颜色是黄色的!")
}

func (y *YellowHuman) Talk() {
	fmt.Println("黄色人种会说话,一般说的都是双字节.")
}

// WhiteHuman 白色人种
type WhiteHuman struct{}

func (w *WhiteHuman) GetColor() {
	fmt.Println("白色人种的皮肤颜色是白色的!")
}

func (w *WhiteHuman) Talk() {
	fmt.Println("白色人种会说话,一般都是单字节.")
}

// AbstractHumanFactory 抽象人类创建工厂
type AbstractHumanFactory interface {
	CreateHuman(t string) Human
}

// HumanFactory 人类创建工厂
type HumanFactory struct{}

// CreateHuman 产生一个人种
func (h *HumanFactory) CreateHuman(t string) Human {
	switch t {
	case "b":
		return &BlackHuman{}
	case "y":
		return &YellowHuman{}
	case "w":
		return &WhiteHuman{}
	default:
		fmt.Println("人种生成错误!")
		return nil
	}
}
