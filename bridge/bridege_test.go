/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/7/7
**/
package bridge

import "testing"

func TestBridge(t *testing.T) {
	c := NewConcreteImplementor()
	refined := NewRefinedAbstraction(c)
	refined.Execute("select * from table")
}
