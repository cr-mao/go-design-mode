/**
User: cr-mao
Date: 2023/11/28 23:17
Email: crmao@qq.com
Desc: singleton2.go
*/
package singleton

import (
	"sync"
)

////声明锁对象
var mutex sync.Mutex

//当对象为空时，对对象加锁，当创建好对象后，获取对象时就不用加锁了
func GetInstance2() *Single {
	if instance == nil {
		mutex.Lock()
		if instance == nil {
			instance = new(Single)
		}
		mutex.Unlock()
	}
	return instance
}
