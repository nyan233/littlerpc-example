package main

import (
	"fmt"
	"github.com/nyan233/littlerpc/impl/client"
	"github.com/nyan233/littlerpc/impl/server"
)

type Hello int

func (receiver Hello) Hello(s string) (int,error) {
	fmt.Println(s)
	return 1 << 20,nil
}

func main() {
	server := server.NewServer(server.WithAddressServer(":1234"))
	err := server.Elem(new(Hello))
	if err != nil {
		panic(err)
	}
	err = server.Start()
	if err != nil {
		panic(err)
	}
	clientInfo := new(Hello)
	client,err := client.NewClient(client.WithAddressClient(":1234"))
	if err != nil {
		panic(err)
	}
	_ = client.BindFunc(clientInfo)
	rep, _ := client.Call("Hello.Hello", "hello world!")
	fmt.Println(rep[0])
}
