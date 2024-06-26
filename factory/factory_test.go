/**
* @Author: maozhongyu
* @Desc:
* @Date: 2024/6/27
**/
package factory

import (
	"testing"
)

func TestFactory(t *testing.T) {
	anta, _ := MakeClothes("Anta")
	peak, _ := MakeClothes("PEAK")
	PrintDetails(anta)
	PrintDetails(peak)

}
