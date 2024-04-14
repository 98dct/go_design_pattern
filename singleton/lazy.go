package singleton

import "sync"

/*
*

	懒汉式非并发安全，需要加锁
	once.Do(f func()) 和 sync.Mutex作用相同，只是更优雅一点
*/
type Lazy struct {
}

var InsLazy *Lazy
var once sync.Once

func GetInsLazy() *Lazy {
	once.Do(func() {
		InsLazy = &Lazy{}
	})

	return InsLazy
}
