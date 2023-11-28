/**
User: cr-mao
Date: 2023/11/28 23:05
Email: crmao@qq.com
Desc: singleton.go
*/
package singleton

import (
	"sync"
)

type Single struct {
	data int
}

var once sync.Once
var instance *Single

// GetInstance 用于获取单例模式对象
func GetInstance() *Single {
	once.Do(func() {
		instance = &Single{}
	})
	return instance
}
