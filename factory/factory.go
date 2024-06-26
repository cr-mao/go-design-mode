/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package factory

import "fmt"

type Iclothes interface {
	SetName(name string)
	SetSize(size int)
	GetName() string
	GetSize() int
}

type clothes struct {
	name string
	size int
}

func (c *clothes) SetName(name string) {
	c.name = name
}

func (c *clothes) SetSize(size int) {
	c.size = size
}

func (c *clothes) GetName() string {
	return c.name
}

func (c *clothes) GetSize() int {
	return c.size
}

type PEAK struct {
	clothes
}

type Anta struct {
	clothes
}

func NewPEAK() *PEAK {
	return &PEAK{
		clothes: clothes{
			name: "PEAK",
			size: 40,
		},
	}
}
func NewAnta() *Anta {
	return &Anta{
		clothes: clothes{
			name: "Anta",
			size: 30,
		},
	}
}

func MakeClothes(name string) (Iclothes, error) {
	switch name {
	case "PEAK":
		return NewPEAK(), nil
	case "Anta":
		return NewAnta(), nil
	default:
		return nil, fmt.Errorf("wrong clothes type passwd ")
	}
}

func PrintDetails(c Iclothes) {
	fmt.Println(c.GetName(), c.GetSize())
}
