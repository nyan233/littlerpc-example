package main

import (
	"fmt"
	"github.com/nyan233/littlerpc/impl/client"
	"github.com/nyan233/littlerpc/impl/server"
)

type Hello struct{}

type UserJson struct {
	Name string
	Id   int64
}

func (h *Hello) Hello(name string, id int64) (*UserJson, error) {
	return &UserJson{
		Name: name,
		Id:   id,
	}, nil
}

func Server() {
	server := server.NewServer(server.WithAddressServer(":1234"))
	err := server.Elem(&Hello{})
	if err != nil {
		panic(err)
	}
	err = server.Start()
	if err != nil {
		panic(err)
	}
}

func Client() {
	c,err := client.NewClient(client.WithAddressClient(":1234"))
	if err != nil {
		panic(err)
	}
	_ = c.BindFunc(&Hello{})
	rep, err := c.Call("Hello.Hello", "Tony", 1<<20)
	if err != nil {
		panic(err)
	}
	user := rep[0].(*UserJson)
	uErr,_ := rep[1].(error)
	fmt.Println(user)
	fmt.Println(uErr)
	if err != nil {
		panic(err)
	}
}

func main() {
	Server()
	Client()
}
