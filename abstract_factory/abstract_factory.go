package abstract_factory

import "fmt"

// AbstractProductA 抽象产品类A
type AbstractProductA interface {
	DoSomethingA()
}

// AbstractProductB 抽象产品类B
type AbstractProductB interface {
	DoSomethingB()
}

// ProductA1 产品A1的实现类
type ProductA1 struct{}

func (p *ProductA1) DoSomethingA() {
	fmt.Println("产品A1的实现方法")
}

// ProductA2 产品A2的实现类
type ProductA2 struct{}

func (p *ProductA2) DoSomethingA() {
	fmt.Println("产品A2的实现方法")
}

// ProductB1 产品B1的实现类
type ProductB1 struct{}

func (p *ProductB1) DoSomethingB() {
	fmt.Println("产品B1的实现方法")
}

// ProductB2 产品B2的实现类
type ProductB2 struct{}

func (p *ProductB2) DoSomethingB() {
	fmt.Println("产品B2的实现方法")
}

// AbstractCreator 抽象工厂类
type AbstractCreator interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// Creator1 产品等级1的工厂实现类
type Creator1 struct{}

func (c *Creator1) CreateProductA() AbstractProductA {
	return &ProductA1{}
}

func (c *Creator1) CreateProductB() AbstractProductB {
	return &ProductB1{}
}

// Creator2 产品等级2的工厂实现类
type Creator2 struct{}

func (c *Creator2) CreateProductA() AbstractProductA {
	return &ProductA2{}
}

func (c *Creator2) CreateProductB() AbstractProductB {
	return &ProductB2{}
}
