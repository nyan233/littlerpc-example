/*
	@Generator   : littlerpc-generator
	@CreateTime  : 2022-06-21 02:29:20.318733 +0800 CST m=+0.000555065
	@Author      : littlerpc-generator
*/
package main

import (
	"github.com/nyan233/littlerpc/impl/client"
)

type HelloServer1Interface interface{ Hello() (string, error) }

type HelloServer1Proxy struct {
	*client.Client
}

func NewHelloServer1Proxy(client *client.Client) HelloServer1Interface {
	proxy := &HelloServer1Proxy{}
	err := client.BindFunc(proxy)
	if err != nil {
		panic(err)
	}
	proxy.Client = client
	return proxy
}

func (proxy HelloServer1Proxy) Hello() (string, error) {
	inter, err := proxy.Call("HelloServer1.Hello")
	if err != nil {
		return "", err
	}
	r0 := inter[0].(string)
	r1, _ := inter[1].(error)
	return r0, r1
}
