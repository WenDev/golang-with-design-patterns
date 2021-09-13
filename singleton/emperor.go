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
