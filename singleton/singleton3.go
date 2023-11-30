/**
User: cr-mao
Date: 2023/11/28 23:22
Email: crmao@qq.com
Desc: singleton3.go
*/
package singleton

var SingleInstance *Single

func init() {
	SingleInstance = &Single{}
}
