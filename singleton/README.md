# Go语言设计模式（1）单例模式

## 单例模式的定义

个人认为单例模式是23种设计模式中最简单也最好理解的一种，定义如下：

> Ensure a class has only one instance, and provide a global pointof access to it.
>
> 确保一个类只有一个实例，而且自行实例化并向整个系统提供这个实例。
>
> ——《设计模式之禅》

那它有什么用呢？我目前在项目中遇到的最多的需要使用单例模式情况就是工具类——工具类一般都没有必要用一次就新建一个实例，所以使用单例模式来实现是非常合适的，当然到目前为止我只在Java中遇到过这个场景（毕竟Kotlin有语法层面的支持（object），Golang则很少需要这么做）。还有就是如果创建一个实例需要很大的资源开销（比如建立数据库连接等），那么也可以考虑使用单例模式。

## 单例模式的简单例子

我们使用Go语言重写《设计模式之禅》使用的臣子和皇帝的例子：

**singleton/emperor.go**

```go
package singleton

import "fmt"

var instance *emperor // 实例

// emperor 皇帝结构体
// 这里不导出是因为如果导出(首字母大写),那么在其他包里就可以用e := &Emperor{}来创建新的实例了
// 那样就不是单例模式了,我们写这么一堆东西也就没有意义了
type emperor struct {
}

func init() {
	// 初始化一个皇帝
	instance = &emperor{}
}

// GetInstance 得到实例
func GetInstance() *emperor {
	return instance
}

// Say 皇帝发话了
func (e *emperor) Say() {
	fmt.Println("我就是皇帝某某某...")
}

```

臣子类改写成了单元测试：

**singleton/emperor_test.go**

```go
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

```

运行结果：

```
=== RUN   TestGetInstance
我就是皇帝某某某...
我就是皇帝某某某...
我就是皇帝某某某...
--- PASS: TestGetInstance (0.00s)
PASS
```

简单分析一下这个例子：为什么皇帝是单例的呢？其原因是`init()`函数仅在包第一次被加载时执行一次，所以只会创建出一个实例，而我们把`emperor`声明为包外不可访问的了，所以在包的外部无法通过`e := &emperor{}`或者`e := new(emperor)`这种方式创建出新的实例，这就实现了自行实例化并且只有一个实例。

## 懒汉式与饿汉式

单例模式的一个常见考点就是“懒汉式”与“饿汉式”。那么在Go语言里如何编写呢？

### 饿汉式

因为饿汉式相对比较好理解一些，代码写起来也更简单，所以我们先讲讲饿汉式。

顾名思义，饿汉很饿，所以它不等你用到实例就先把实例先给创建好了。这种方法不需要加锁，没有线程安全问题，但是会减慢启动速度，且由于在使用之前就创建了实例，所以会浪费一部分内存空间（也就是说不是“按需创建”）。这种方法适用于创建实例使用的资源比较少的场景。

实际上，我们刚刚写的皇帝与臣子的代码就是饿汉式写法的一个例子（使用`init()`函数）。下面给出饿汉式的通用代码：

写法1：

**singleton/singleton_hungry.go**

```go
package singleton

// 饿汉式写法1: 使用全局变量

var instance1 = &singleton1{}

type singleton1 struct{}

func GetInstance1() *singleton1 {
	return instance1
}

// 饿汉式写法2: 使用init函数

var instance2 *singleton2

type singleton2 struct{}

func init() {
	instance2 = &singleton2{}
}

func GetInstance2() *singleton2 {
	return instance2
}

```

需要注意的是两种写法使用起来差不多，因为虽然全局变量的初始化会比`init()`函数执行早一点，但都是在`main()`函数之前，所以在使用上没有特别大的差距，具体选择哪种还是要看实际的业务场景。

### 懒汉式

有饿汉式自然就会有懒汉式。懒汉式本质上就是按需创建，在你需要用到这个实例的时候才会去创建它。这种方法写起来比较复杂（但也有使用`sync.Once`的简单写法），可能会产生线程安全问题，适用于创建实例使用的资源较多的场景。

懒汉式有很多种写法，它们是否线程安全也是不一样的，下面来介绍一下这些写法（注：以下所有代码都在**singleton/singleton_lazy.go**文件中）：

#### 写法1：不加锁

```go
// 写法1: 不加锁, 线程不安全

var instance3 *singleton3

type singleton3 struct{}

func GetInstance3() *singleton3 {
	if instance3 == nil {
		instance3 = &singleton3{}
	}
	return instance3
}
```

相信大家都能看出来这种方法是线程不安全的，在并发执行的时候可能会由于多个线程同时判断`instance3 == nil`成立进而创建多个实例，所以不推荐使用。

#### 写法2：对GetInstance()方法加锁

```go
// 写法2: 对GetInstance方法加锁, 线程安全, 但是效率比较低

var (
	instance4 *singleton4
	lock1     sync.Mutex
)

type singleton4 struct{}

func GetInstance4() *singleton4 {
	lock1.Lock()
	defer lock1.Unlock()
	if instance4 == nil {
		instance4 = &singleton4{}
	}
	return instance4
}

```

由于在多线程并发时`GetInstance4()`方法只允许一个线程进入，第二个线程需要在第一个线程退出之后才能进入，所以这种方法是线程安全的。但是它也有显而易见的缺点：效率低，因为每次获取实例时都需要加锁解锁。

#### 写法3：创建单例时加锁

```go
// 写法3: 创建单例时加锁, 线程不安全, 这种写法仅仅是为了引出写法4

var (
	instance5 *singleton5
	lock2     sync.Mutex
)

type singleton5 struct{}

func GetInstance5() *singleton5 {
	if instance5 == nil {
		lock2.Lock()
		instance5 = &singleton5{}
		lock2.Unlock()
	}
	return instance5
}

```

这种方法也是线程不安全的。虽然同一时刻只可能有一个线程在执行`instance5 = &singleton5{}`这行代码，但是仍然有可能有多个线程都判断`instance5 == nil`成立并创建多个对象。它本质上跟不加锁没什么区别，提及这种写法仅仅是为了引出下面的写法4：双重检查机制。

#### 写法4：双重检查

```go
// 写法4: 双重检查, 线程安全

var (
	instance6 *singleton6
	lock3     sync.Mutex
)

type singleton6 struct{}

func GetInstance6() *singleton6 {
	if instance6 == nil {
		lock3.Lock()
		if instance6 == nil {
			instance6 = &singleton6{}
		}
		lock3.Unlock()
	}
	return instance6
}

```

这是一种线程安全的写法。既然有可能有多个线程同时判断`instance6 == nil`，那么再加锁之后再检查一次就行了。但是每一次获取实例都要加锁还要检查两次显然不是一个明智的选择，所以我们有更优的解法：使用`sync.Once`。

#### 写法5：使用sync.Once

```go
// 写法5: 使用sync.Once, 线程安全, 推荐使用

var (
	instance7 *singleton7
	once      sync.Once
)

type singleton7 struct{}

func GetInstance7() *singleton7 {
	once.Do(func() {
		instance7 = &singleton7{}
	})
	return instance7
}

```

`sync.Once`有点类似于`init()`函数，它们都执行且仅执行一次，区别在于`sync.Once`是在你需要的时候执行，而`init()`是在包第一次被加载的时候执行。那为什么`sync.Once`可以解决加锁的问题呢？这就跟`sync.Once`的内部实现有关了。

以下是`sync.Once`的源码，非常短，但是很有参考价值：

```go
type Once struct {
	done uint32
	m    Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

```

可以发现`Do()`函数中仅仅做了一次判断——如果传入的函数已经执行了（`done`的值为1），那么就不执行，直接返回；否则执行`doSlow()`方法。在`doSlow()`方法中进行了加锁并执行了传入的函数，在代码块运行结束后再把`done`修改为1，这样就实现了执行且仅执行一次的功能，并且只有第一次需要加锁，这样对于`GetInstance()`函数来说就不再需要判断`instance`是否为`nil`了，也不再需要手动进行加锁解锁操作了，可谓是非常棒的一种解决方案。

## 总结

Go语言实现单例模式还是挺简单的，基本上看一遍就能懂（从Java转到Go的我表示：比的Java简单多了！尤其是sync.Once写法，精彩程度堪比Java单例模式的enum写法），但要注意转变思维——因为Go语言本身的特点，它的单例模式写法与其他语言（Java、C++等）有很大的区别，如果是初学者自然不用在意这个，但是对于有其他语言基础的还是应该注意一下。