/*
单例模式确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例，这个类称为单例类，它提供全局访问的方法。
单例模式的要点有三个：一是某个类只能有一个实例；二是它必须自行创建这个实例；三是它必须自行向整个系统提供这个实例。
*/

/*type Once struct {
     done uint32    //执行标记
     m    Mutex    //互斥锁，用于在代码上创建一个临界区,保证同一时间只有一个goroutine可以执行这个临界区代码。只有两个公开方法：Lock()加锁，unlock()解锁。
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {    //原子操作判断o.done是否为1，若为1，表示f已经执行过，直接返回
		o.m.Lock()   //加锁，保证互斥访问
		defer o.m.Unlock()
		if o.done == 0 {    /基o.done为0，表示f未执行，对o.done原子赋值，并且执行f函数
			defer atomic.StoreUint32(&o.done, 1)
			f()
        }
    }
}*/

package singleton

import "sync"

type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}