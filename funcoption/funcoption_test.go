/**
User: cr-mao
Date: 2023/11/30 22:29
Email: crmao@qq.com
Desc: function_option.go
*/
package funcoption

import (
	"fmt"
	"testing"
)

type DbConfig struct {
	host     string
	port     int
	password string
	username string
	printLog bool
}

type Option func(config *DbConfig)

func WithPrintLog() Option {
	return func(config *DbConfig) {
		config.printLog = true
	}
}

func NewDbConfig(Options ...Option) *DbConfig {
	//默认值
	var config = &DbConfig{
		host: "127.0.0.1",
		port: 3306,
	}
	for _, option := range Options {
		option(config)
	}
	return config
}

func TestFuncOption(t *testing.T) {
	dbConfig := NewDbConfig(WithPrintLog())
	fmt.Printf("%+v\n", dbConfig) // &{host:127.0.0.1 port:3306 password: username: printLog:true}
}
