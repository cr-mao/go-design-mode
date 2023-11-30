/**
User: cr-mao
Date: 2023/11/28 23:08
Email: crmao@qq.com
Desc: singleton_test.go
*/
package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	instances := [100]*Single{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	wg.Wait()
	for i := 1; i < 100; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}

func TestMutexSingle(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	instances := [100]*Single{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			instances[index] = GetInstance2()
			wg.Done()
		}(i)
	}
	// 关闭channel，所有协程同时开始运行，实现并行(parallel)
	wg.Wait()
	for i := 1; i < 100; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
