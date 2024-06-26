/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package objectpool

import (
	"fmt"
	"testing"
)

func TestObjectPool(t *testing.T) {
	num := func() interface{} {
		return 10.0
	}
	pool := NewPool(num)
	object := pool.Acquire()

	fmt.Println(pool.Inuse)
	fmt.Println(pool.Available)

	pool.Release(object)
	fmt.Println(pool.Inuse)
	fmt.Println(pool.Available)
}
