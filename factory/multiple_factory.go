package factory

import "fmt"

// Human1 人类总称
type Human1 interface {
	GetColor() // 每个人种的皮肤都有相应的颜色
	Talk()     // 人类会说话
}

// BlackHuman1 黑色人种
type BlackHuman1 struct{}

func (b *BlackHuman1) GetColor() {
	fmt.Println("黑色人种的皮肤颜色是黑色的!")
}

func (b *BlackHuman1) Talk() {
	fmt.Println("黑人会说话,一般人听不懂.")
}

// YellowHuman1 黄色人种
type YellowHuman1 struct{}

func (y *YellowHuman1) GetColor() {
	fmt.Println("黄色人种的皮肤颜色是黄色的!")
}

func (y *YellowHuman1) Talk() {
	fmt.Println("黄色人种会说话,一般说的都是双字节.")
}

// WhiteHuman1 白色人种
type WhiteHuman1 struct{}

func (w *WhiteHuman1) GetColor() {
	fmt.Println("白色人种的皮肤颜色是白色的!")
}

func (w *WhiteHuman1) Talk() {
	fmt.Println("白色人种会说话,一般都是单字节.")
}

// AbstractHumanFactory1 抽象人类创建工厂
type AbstractHumanFactory1 interface {
	// CreateHuman 产生一个人种
	// 注意,由于每个子类都有与其自身关联的产品类,所以这里已经不需要通过传递参数来指明创建哪种产品了
	CreateHuman() Human
}

// BlackHumanFactory 黑色人种创建工厂
type BlackHumanFactory struct{}

// CreateHuman 产生一个黑色人种
func (h *BlackHumanFactory) CreateHuman() Human {
	return &BlackHuman{}
}

// YellowHumanFactory 黄色人种创建工厂
type YellowHumanFactory struct{}

// CreateHuman 产生一个黄色人种
func (h *YellowHumanFactory) CreateHuman() Human {
	return &YellowHuman{}
}

// WhiteHumanFactory 白色人种创建工厂
type WhiteHumanFactory struct{}

// CreateHuman 产生一个白色人种
func (h *WhiteHumanFactory) CreateHuman() Human {
	return &WhiteHuman{}
}
