/*
单例模式：单例模式是一种常用的软件设计模式。在它的核心结构中只包含一个被称为单例类的特殊类。
通过单例模式可以保证系统中一个类只有一个实例而且该实例易于外界访问，从而方便对实例个数的控制并节约系统资源。
如果希望在系统中某个类的对象只能存在一个，单例模式是最好的解决方案。
*/
package single

import "sync"

type singleton map[string]string

var (
	once     sync.Once //只执行一次
	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
