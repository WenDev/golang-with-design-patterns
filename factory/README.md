# Go语言设计模式（2）工厂模式

## 工厂模式的定义

工厂模式的定义如下：

> Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.
>
> 定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类。
>
> ——《设计模式之禅》

工厂模式在Go语言中还是很常见的。因为Go语言的结构体没有构造器，而工厂模式本质上就是new一个对象的替代品，所以在项目中一般都需要使用简单工厂模式去实例化一个结构体已达到与构造器相同的目的（即我们经常见到的`NewXxx()`方法，关于什么是简单工厂模式我们马上就会讲到）；另外在DDD（领域驱动设计）中也经常可以看到工厂模式的身影。

## 工厂模式的简单例子

我们使用Go语言重写《设计模式之禅》中给出的女娲造人的例子：

**factory/human_factory.go**

```go
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

func (w *WhiteHuman) Talk()  {
	fmt.Println("白色人种会说话,一般都是单字节.")
}

// AbstractHumanFactory 抽象人类创建工厂
type AbstractHumanFactory interface {
	CreateHuman(t string) Human
}

// HumanFactory 人类创建工厂
type HumanFactory struct {}

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

```

**factory/human_factory_test.go**

```go
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

```

运行结果如下：

```
=== RUN   TestHumanFactory_CreateHuman
--造出的第一批人是白色人种--
白色人种的皮肤颜色是白色的!
白色人种会说话,一般都是单字节.
--造出的第二批人是黑色人种--
黑色人种的皮肤颜色是黑色的!
黑人会说话,一般人听不懂.
--造出的第三批人是黄色人种--
黄色人种的皮肤颜色是黄色的!
黄色人种会说话,一般说的都是双字节.
--- PASS: TestHumanFactory_CreateHuman (0.00s)
PASS
```

与Java版本的代码类似，Go语言版本的代码也非常简单，有过Go语言项目开发经历的话应该不难理解。

## 工厂模式的通用代码

**factory/factory.go**

```go
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

```

实际使用效果如下：

**factory/factory_test.go**

```go
package factory

import "testing"

func TestConcreteCreator_CreateProduct(t *testing.T) {
	c := &ConcreteCreator{}
	c1 := c.CreateProduct("1")
	c1.Method()
	c2 := c.CreateProduct("2")
	c2.Method()
}

```

运行结果：

```
=== RUN   TestConcreteCreator_CreateProduct
我是具体产品类1的业务逻辑
我是具体产品类2的业务逻辑
--- PASS: TestConcreteCreator_CreateProduct (0.00s)
PASS
```

## 工厂模式的扩展

除了前面提到的工厂模式之外，工厂模式还有以下扩展：

### 简单工厂模式

如果只有一个工厂类，就可以将抽象工厂类去掉，只保留实现类，此时我们就可以发现我们没有必要将工厂类实例化，那么我们就将工厂类的方法设置为静态（在Go语言中就是改成不属于任何结构体的函数）。

这就是简单工厂模式，它是工厂模式的一种弱化，虽然不便于扩展但是很常用，例如我们在项目中用来实例化结构体的`NewXxx()`函数就是简单工厂模式在Go语言中最简单也最常见应用。

还是以女娲造人为例，使用简单工厂模式实现的代码如下：

**factory/simple_factory.go**

```go
package factory

import "fmt"

func CreateHuman(t string) Human {
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

```

对工厂类的改动同时也引起了调用方的改动——不需要将工厂类实例化了，直接调用`CreateHuman()`即可。

**factory/simple_factory_test.go**

```go
package factory

import (
	"fmt"
	"testing"
)

func TestCreateHuman(t *testing.T) {
	// 女娲第一次造人,火候不足,于是白人产生了
	fmt.Println("--造出的第一批人是白色人种--")
	wh := CreateHuman("w")
	wh.GetColor()
	wh.Talk()
	// 女娲第二次造人,火候过足,于是黑人产生了
	fmt.Println("--造出的第二批人是黑色人种--")
	bh := CreateHuman("b")
	bh.GetColor()
	bh.Talk()
	// 第三次造人,火候刚刚好,于是黄色人种产生了
	fmt.Println("--造出的第三批人是黄色人种--")
	yh := CreateHuman("y")
	yh.GetColor()
	yh.Talk()
}

```

运行结果：

```
=== RUN   TestCreateHuman
--造出的第一批人是白色人种--
白色人种的皮肤颜色是白色的!
白色人种会说话,一般都是单字节.
--造出的第二批人是黑色人种--
黑色人种的皮肤颜色是黑色的!
黑人会说话,一般人听不懂.
--造出的第三批人是黄色人种--
黄色人种的皮肤颜色是黄色的!
黄色人种会说话,一般说的都是双字节.
--- PASS: TestCreateHuman (0.00s)
PASS
```

### 多工厂类的工厂模式

随着项目的不断扩展，实例化一个对象可能非常耗费精力，或者工厂类需要实例化很多种不同的对象，这时工厂方法就会变得很长，导致难以阅读。我们可以让工厂类也继承于一个基类，每一个产品类对应一个工厂类，这时工厂类的职责也就变得更加明确了：生产与自己关联的产品。虽然这带来了维护和扩展上的挑战——因为需要同时维护产品类及对应于该产品类的工厂类，但是这样的设计也降低了工厂类的复杂度。

以女娲造人为例，多个工厂类的工厂模式代码如下：

**factory/multiple_factory.go**

```go
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

// WhiteHuman 白色人种
type WhiteHuman1 struct{}

func (w *WhiteHuman1) GetColor() {
	fmt.Println("白色人种的皮肤颜色是白色的!")
}

func (w *WhiteHuman1) Talk()  {
	fmt.Println("白色人种会说话,一般都是单字节.")
}

// AbstractHumanFactory1 抽象人类创建工厂
type AbstractHumanFactory1 interface {
	// CreateHuman 产生一个人种
	// 注意,由于每个子类都有与其自身关联的产品类,所以这里已经不需要通过传递参数来指明创建哪种产品了
	CreateHuman() Human
}

// BlackHumanFactory 黑色人种创建工厂
type BlackHumanFactory struct {}

// CreateHuman 产生一个黑色人种
func (h *BlackHumanFactory) CreateHuman() Human {
	return &BlackHuman{}
}

// YellowHumanFactory 黄色人种创建工厂
type YellowHumanFactory struct {}

// CreateHuman 产生一个黄色人种
func (h *YellowHumanFactory) CreateHuman() Human {
	return &YellowHuman{}
}

// WhiteHumanFactory 白色人种创建工厂
type WhiteHumanFactory struct {}

// CreateHuman 产生一个白色人种
func (h *WhiteHumanFactory) CreateHuman() Human {
	return &WhiteHuman{}
}

```

调用方也需要一定的更改——在实例化每种产品之前先实例化对应的工厂：

**factory/multiple_factory_test.go**

```go
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

```

运行结果：

```
=== RUN   TestCreateHuman1
--造出的第一批人是白色人种--
白色人种的皮肤颜色是白色的!
白色人种会说话,一般都是单字节.
--造出的第二批人是黑色人种--
黑色人种的皮肤颜色是黑色的!
黑人会说话,一般人听不懂.
--造出的第三批人是黄色人种--
黄色人种的皮肤颜色是黄色的!
黄色人种会说话,一般说的都是双字节.
--- PASS: TestCreateHuman1 (0.00s)
PASS
```

## 总结

个人认为工厂模式是在Go语言中应用最广泛的一种设计模式——因为Go语言的结构体没有构造器，所以在大型项目中或遇到创建结构体的过程特别复杂的场景时，需要使用简单工厂模式来帮助我们创建结构体。另外，在领域驱动设计（`DDD`，Domain Driven Design）中，创建聚合根等场景中也用到了工厂模式，尤其是当创建聚合根的逻辑不适合放在聚合根上时。


