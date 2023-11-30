/**
User: cr-mao
Date: 2023/11/30 22:20
Email: crmao@qq.com
Desc: decorator_test.go
*/
package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	//具体零件
	phone := &BaseParts{}
	fmt.Printf("基础零件的价格为：%f\n", phone.GetPrice())

	//定义添加IPhone手机
	iPhone := &IPhone{}
	iPhone.SetComponent(phone)
	fmt.Printf("苹果的价格为：%f\n", iPhone.GetPrice())

	//定义添加Xiaomi手机
	xiaomi := &Xiaomi{}
	xiaomi.SetComponent(phone)
	fmt.Printf("小米的价格为：%f\n", xiaomi.GetPrice())

}
