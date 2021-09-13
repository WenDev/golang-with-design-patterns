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
