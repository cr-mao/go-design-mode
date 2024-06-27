/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package component

import "fmt"

// 组件接口
type Component interface {
	Execute()
}

type Leaf struct {
	name string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

func (l *Leaf) Execute() {
	fmt.Println(l.name)
}

// 组合类
type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{
		children: make([]Component, 0),
	}
}

// 添加子组件到组合中
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

// 遍历子组件
func (c *Composite) Execute() {
	for _, child := range c.children {
		child.Execute()
	}
}
