package factory

import "fmt"

// Product 抽象产品类
type Product interface {
	// Method 抽象方法
	Method()
}

// ConcreteProduct1 具体产品类1
type ConcreteProduct1 struct{}

func (c *ConcreteProduct1) Method() {
	fmt.Println("我是具体产品类1的业务逻辑")
}

// ConcreteProduct2 具体产品类2
type ConcreteProduct2 struct{}

func (c *ConcreteProduct2) Method() {
	fmt.Println("我是具体产品类2的业务逻辑")
}

// Creator 抽象工厂类
// 创建一个产品对象,其输入参数类型可以自行设置,这里用了string
type Creator interface {
	CreateProduct(t string) Product
}

// ConcreteCreator 具体工厂类
type ConcreteCreator struct{}

func (c *ConcreteCreator) CreateProduct(t string) Product {
	switch t {
	case "1":
		return &ConcreteProduct1{}
	case "2":
		return &ConcreteProduct2{}
	default:
		// 异常处理
		return nil
	}
}
