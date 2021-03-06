package main

import (
	"fmt"
	"github.com/nyan233/littlerpc/impl/client"
	"github.com/nyan233/littlerpc/impl/server"
	"math"
	"time"
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
	server := server.NewServer(server.WithAddressServer("127.0.0.1:8080","127.0.0.1:9090"))
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
	c,err := client.NewClient(client.WithBalance("roundRobin"))
	if err != nil {
		panic(err)
	}
	_ = c.BindFunc(&Hello{})
	rep, err := c.Call("Hello.Hello", "Tony", 1<<20)
	if err != nil {
		panic(err)
	}
	user := rep[0].(*UserJson)
	fmt.Println(user)
	if err != nil {
		panic(err)
	}
}

func main() {
	Server()
	// 根据规则开启负载均衡
	err := client.OpenBalance("live", "live://127.0.0.1:8080;127.0.0.1:9090",
		time.Duration(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	Client()
	Client()
}
