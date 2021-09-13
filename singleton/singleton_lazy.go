package singleton

import "sync"

// 写法1: 不加锁, 线程不安全

var instance3 *singleton3

type singleton3 struct{}

func GetInstance3() *singleton3 {
	if instance3 == nil {
		instance3 = &singleton3{}
	}
	return instance3
}

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
