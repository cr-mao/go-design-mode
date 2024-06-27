/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package component

import "testing"

func TestComponent(t *testing.T) {
	component1 := NewComposite()
	leaf1 := NewLeaf("leaf1")
	leaf2 := NewLeaf("leaf2")
	component1.Add(leaf1)
	component1.Add(leaf2)
	component2 := NewComposite()
	leaf3 := NewLeaf("leaf3")
	component2.Add(leaf3)
	component1.Add(component2)
	component1.Execute()
	/*
		leaf1
		leaf2
		leaf3
	*/
}
