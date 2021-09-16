# Go语言设计模式（3）抽象工厂模式

## 抽象工厂模式的定义

抽象工厂模式的定义如下：

> Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
>
> 为创建一组相关或者相互依赖的对象提供一个接口，而且无须指定它们的具体类。
>
> ——《设计模式之禅》

好吧，就跟它的名字一样，它的定义也很抽象，很让人摸不着头脑。但其实这种设计模式还是很好理解的，只要明白了它和工厂模式有什么区别，再看一些实例代码，理解这种设计模式并不是难事。我们接下来就来具体讲讲它。

## 抽象工厂模式和工厂模式的区别

首先考虑一个问题：抽象工厂模式和工厂模式的区别在哪？

抽象工厂模式最大的特征就是定义中提到的“相关或相互依赖的对象”，通俗点说，就是抽象工厂模式相比工厂模式多了两个概念：“**产品族**”和“**等级结构**”。工厂模式中的每个工厂只能生产一种产品，而抽象工厂模式将某些相关、等级结构不同的产品组成了一个“产品族”，然后用一个工厂来统一生产。

还是觉得很难理解？好的，Show you code！

## 抽象工厂模式的简单例子

这里我没有使用《设计模式之禅》的例子，因为它比较复杂，我们来举个更简单更好理解的例子：

我们可以把Web应用简单地看做是由前端（客户端）和后端（服务端）组成的，不知道在看我的文章的同学中有没有知道“前端架构师”这个职业的同学，我们一直以为做后端开发更容易成为架构师，但是当前端变得越来越复杂时，前端架构师这种职位也就产生了。所以，在这个例子中，我们把架构师分为“前端架构师”和“后端架构师”两类。当然，要开发一个完整的Web应用，光有架构师还不行，还得有程序员来把架构师给出的设计用代码实现，而我们都知道程序员有做前端开发的，也有做后端开发的，这样也可以把程序员分为“前端程序员”和“后端程序员”两类。

在这个场景中，我们有四种产品：前端程序员、前端架构师、后端程序员、后端架构师。我们可以把“前端”和“后端”分为两个不同的**等级结构**，我们可以为这两个等级结构创建两个工厂类`FrontEndFactory`和`BackEndFactory`，我们同时也可以把“程序员”和“架构师”分为另外两个不同的**产品族**，在每个工厂类里添加`CreateProgrammer()`和`CreateArchitect()`两个方法用于创建程序员和架构师的实例——也就是说，**有多少个产品族，在抽象工厂类里就有多少个创建方法**。

根据这个场景我们可以画出下面这张UML类图（稍微有点不太标准，不过能通过这张图把上面描述的场景搞懂就OK了）：

![](https://gitee.com/QNKCDZ0/pictures/raw/master/img/20210917005502.png)

接下来用代码实现它：

**abstract_factory/programmer_factory.go**

```go
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

```

代码比较长，可以对照着类图来看。

写个场景类（测试）来调用一下：

**abstract_factory/programmer_factory_test.go**

```go
package abstract_factory

import (
	"fmt"
	"testing"
)

func TestCreateProgrammerAndArchitect(t *testing.T) {
	fmt.Println("前端组招到了一个程序员和一个架构师")
	ff := FrontEndFactory{}
	fa := ff.CreateArchitect()
	fp := ff.CreateProgrammer()
	fmt.Println("前端组接到任务,开始工作...")
	fa.Design()
	fp.Work()
	fmt.Println("后端组招到了一个程序员和一个架构师")
	bf := BackEndFactory{}
	ba := bf.CreateArchitect()
	bp := bf.CreateProgrammer()
	fmt.Println("后端组接到任务,开始工作...")
	ba.Design()
	bp.Work()
}

```

运行结果：

```
=== RUN   TestCreateProgrammerAndArchitect
前端组招到了一个程序员和一个架构师
前端组接到任务,开始工作...
前端架构师做了一个页面秒开的设计
前端程序员在用WebStorm写TypeScript代码
后端组招到了一个程序员和一个架构师
后端组接到任务,开始工作...
后端架构师做了一个可以抗住上万并发的设计
后端程序员在用GoLand写Golang代码
--- PASS: TestCreateProgrammerAndArchitect (0.00s)
PASS
```

看完这个例子之后，是不是彻底明白抽象工厂模式是怎么一回事了呢。

我们也可以看出，在抽象工厂模式中，添加等级结构很方便——假如我们要添加一个新的等级结构“基础架构（BasicArchitecture）”，那么直接新建对应的程序员类`BasicArchitectureProgrammer`和架构师类`BasicArchitectureArchitect`继承`Programmer`和`Architect`，再新建一个工厂类`BasicArchitectureFactory`继承`AbstractFactory`就可以完成扩展。

但是，我们也很容易发现添加产品族很难，且违背了开闭原则——假如我们要添加一个新的工种“测试工程师（TestEngineer）”，则需要建立新的抽象类`TestEngineer`（成员方法可以叫`Test()`），并且在所有的工厂类里面都添加、实现`CreateTestEngineer()`用于创建不同等级结构（前端、后端）的测试工程师，这样才可以完成扩展。这种方法显而易见地违背了开闭原则，增大了维护难度，所以在使用抽象工厂模式时最好一开始就把所有的产品族设计好，尽量减少对产品族的添加和删除。

## 抽象工厂模式的通用代码

**abstract_factory/abstract_factory.go**

```go
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

```

**abstract_factory/abstract_factory_test.go**

```go
package abstract_factory

import "testing"

func TestAbstractFactory(t *testing.T) {
	c1 := Creator1{}
	c2 := Creator2{}
	a1 := c1.CreateProductA()
	a2 := c2.CreateProductA()
	b1 := c1.CreateProductB()
	b2 := c2.CreateProductB()
	a1.DoSomethingA()
	a2.DoSomethingA()
	b1.DoSomethingB()
	b2.DoSomethingB()
}

```

运行结果：

```
=== RUN   TestAbstractFactory
产品A1的实现方法
产品A2的实现方法
产品B1的实现方法
产品B2的实现方法
--- PASS: TestAbstractFactory (0.00s)
PASS
```

## 总结

抽象工厂模式看似很难理解，实际上只要理解了“产品族”和“等级结构”这两个概念，抽象工厂模式就非常简单了。那么它有什么用呢？一个很典型的用途就是适配不同数据库——不同的数据库所提供的操作应该是相同的（增、删、改、查、事务等），但是每种数据库的底层实现又不一样，这时就可以使用抽象工厂模式来生成不同数据库的操作对象；另外，Java的AWT也运用了抽象工厂模式来实现不同操作系统下应用程序界面的统一。