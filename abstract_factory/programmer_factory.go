package abstract_factory

import "fmt"

// Programmer 程序员总称
type Programmer interface {
	Work() // 程序员都会工作
}

// Architect 架构师总称
type Architect interface {
	Design() // 架构师都会做架构设计
}

// FrontEndArchitect 前端架构师
type FrontEndArchitect struct{}

func (a *FrontEndArchitect) Design() {
	fmt.Println("前端架构师做了一个页面秒开的设计")
}

// FrontEndProgrammer 前端程序员
type FrontEndProgrammer struct{}

func (p *FrontEndProgrammer) Work() {
	fmt.Println("前端程序员在用WebStorm写TypeScript代码")
}

// BackEndArchitect 后端架构师
type BackEndArchitect struct{}

func (a *BackEndArchitect) Design() {
	fmt.Println("后端架构师做了一个可以抗住上万并发的设计")
}

// BackEndProgrammer 后端程序员
type BackEndProgrammer struct{}

func (p *BackEndProgrammer) Work() {
	fmt.Println("后端程序员在用GoLand写Golang代码")
}

// AbstractFactory 抽象工厂
type AbstractFactory interface {
	CreateProgrammer() Programmer // 创建程序员
	CreateArchitect() Architect   // 创建架构师
}

// FrontEndFactory 前端工厂
type FrontEndFactory struct{}

func (f *FrontEndFactory) CreateProgrammer() Programmer {
	return &FrontEndProgrammer{}
}

func (f *FrontEndFactory) CreateArchitect() Architect {
	return &FrontEndArchitect{}
}

// BackEndFactory 后端工厂
type BackEndFactory struct{}

func (f *BackEndFactory) CreateProgrammer() Programmer {
	return &BackEndProgrammer{}
}

func (f *BackEndFactory) CreateArchitect() Architect {
	return &BackEndArchitect{}
}
