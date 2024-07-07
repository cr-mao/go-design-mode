/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/7/7
**/
package bridge

import "fmt"

type Implementor interface {
	Query(sql string)
}

type ConcreteImplementor struct {
}

func (*ConcreteImplementor) Query(sql string) {
	fmt.Println("打印信息" + sql)
}

func NewConcreteImplementor() *ConcreteImplementor {
	return &ConcreteImplementor{}
}

// 扩充抽象类

type RefinedAbstraction struct {
	method Implementor
}

func (c *RefinedAbstraction) Execute(sql string) {
	c.method.Query(sql)
}

func NewRefinedAbstraction(im Implementor) *RefinedAbstraction {
	return &RefinedAbstraction{method: im}
}
