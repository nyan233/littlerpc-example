/*
	@Generator   : littlerpc-generator
	@CreateTime  : 2022-06-21 02:33:45.094649 +0800 CST m=+0.000846871
	@Author      : littlerpc-generator
*/
package main

import (
	"github.com/nyan233/littlerpc/impl/client"
)

type FileServerInterface interface {
	SendFile(path string, data []byte) error
	GetFile(path string) ([]byte, bool, error)
	OpenSysFile(path string) ([]byte, error)
}

type FileServerProxy struct {
	*client.Client
}

func NewFileServerProxy(client *client.Client) FileServerInterface {
	proxy := &FileServerProxy{}
	err := client.BindFunc(proxy)
	if err != nil {
		panic(err)
	}
	proxy.Client = client
	return proxy
}

func (proxy FileServerProxy) SendFile(path string, data []byte) error {
	inter, err := proxy.Call("FileServer.SendFile", path, data)
	if err != nil {
		return err
	}
	r0, _ := inter[0].(error)
	return r0
}

func (proxy FileServerProxy) GetFile(path string) ([]byte, bool, error) {
	inter, err := proxy.Call("FileServer.GetFile", path)
	if err != nil {
		return nil, false, err
	}
	r0 := inter[0].([]byte)
	r1 := inter[1].(bool)
	r2, _ := inter[2].(error)
	return r0, r1, r2
}

func (proxy FileServerProxy) OpenSysFile(path string) ([]byte, error) {
	inter, err := proxy.Call("FileServer.OpenSysFile", path)
	if err != nil {
		return nil, err
	}
	r0 := inter[0].([]byte)
	r1, _ := inter[1].(error)
	return r0, r1
}
