/*
	@Generator   : littlerpc-generator
	@CreateTime  : 2022-06-21 02:46:52.594675 +0800 CST m=+0.000821246
	@Author      : littlerpc-generator
*/
package main

import (
	"github.com/nyan233/littlerpc/impl/client"
)

type HelloInterface interface {
	SayHelloToProtoBuf(pb *Student) (*Student, error)
	SayHelloToJson(jn *Student) (*Student, error)
}

type HelloProxy struct {
	*client.Client
}

func NewHelloProxy(client *client.Client) HelloInterface {
	proxy := &HelloProxy{}
	err := client.BindFunc(proxy)
	if err != nil {
		panic(err)
	}
	proxy.Client = client
	return proxy
}

func (proxy HelloProxy) SayHelloToProtoBuf(pb *Student) (*Student, error) {
	inter, err := proxy.Call("Hello.SayHelloToProtoBuf", pb)
	if err != nil {
		return nil, err
	}
	r0 := inter[0].(*Student)
	r1, _ := inter[1].(error)
	return r0, r1
}

func (proxy HelloProxy) SayHelloToJson(jn *Student) (*Student, error) {
	inter, err := proxy.Call("Hello.SayHelloToJson", jn)
	if err != nil {
		return nil, err
	}
	r0 := inter[0].(*Student)
	r1, _ := inter[1].(error)
	return r0, r1
}
