/**
* @Author: maozhongyu
* @Desc: 外观模式
* @Date: 2024/7/10
**/
package facade

import "fmt"

type Facade struct {
	subsystemA SubsystemA
	subsystemB SubsystemB
}

type SubsystemA struct {
}
type SubsystemB struct {
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: SubsystemA{},
		subsystemB: SubsystemB{},
	}
}

func (c *Facade) MethodA() {
	c.subsystemB.MethodFour()
	c.subsystemA.MethodOne()
	c.subsystemB.MethodThree()
}
func (c *Facade) MethodB() {
	c.subsystemA.MethodTwo()
	c.subsystemB.MethodFour()
}

func (c *SubsystemA) MethodOne() {
	fmt.Println("SubsystemA MethodOne")
}

func (c *SubsystemA) MethodTwo() {
	fmt.Println("SubsystemA Methodtwo")
}

func (c *SubsystemB) MethodThree() {
	fmt.Println("SubsystemA MethodThree")
}

func (c *SubsystemB) MethodFour() {
	fmt.Println("SubsystemA MethodFour")
}
