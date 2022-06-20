/*
	@Generator   : littlerpc-generator
	@CreateTime  : 2022-06-21 02:29:16.138064 +0800 CST m=+0.000663063
	@Author      : littlerpc-generator
*/
package main

import (
	"github.com/nyan233/littlerpc/impl/client"
)

type HelloServer2Interface interface {
	Init(str string) error
	Hello() (string, error)
}

type HelloServer2Proxy struct {
	*client.Client
}

func NewHelloServer2Proxy(client *client.Client) HelloServer2Interface {
	proxy := &HelloServer2Proxy{}
	err := client.BindFunc(proxy)
	if err != nil {
		panic(err)
	}
	proxy.Client = client
	return proxy
}

func (proxy HelloServer2Proxy) Init(str string) error {
	inter, err := proxy.Call("HelloServer2.Init", str)
	if err != nil {
		return err
	}
	r0, _ := inter[0].(error)
	return r0
}

func (proxy HelloServer2Proxy) Hello() (string, error) {
	inter, err := proxy.Call("HelloServer2.Hello")
	if err != nil {
		return "", err
	}
	r0 := inter[0].(string)
	r1, _ := inter[1].(error)
	return r0, r1
}
