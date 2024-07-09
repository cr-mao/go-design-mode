/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/7/10
**/
package facade

import (
	"fmt"
	"testing"
)

func TestNewFacade(t *testing.T) {
	facade := NewFacade()
	facade.MethodA()
	fmt.Println("----------")
	facade.MethodB()
}
