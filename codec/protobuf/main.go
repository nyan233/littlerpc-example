package main

import (
	"fmt"
	"github.com/nyan233/littlerpc/impl/client"
	"github.com/nyan233/littlerpc/impl/server"
	"github.com/nyan233/littlerpc/protocol"
)



type Hello struct {}

func (h *Hello) SayHelloToProtoBuf(pb *Student) (*Student,error) {
	fmt.Println(pb)
	return &Student{
		Name:          "Jenkins",
		Male:          false,
		Scores:        []int32{2,4,8,16,32,64,128},
	},nil
}

func (h *Hello) SayHelloToJson(jn *Student) (*Student,error) {
	fmt.Println(jn)
	return &Student{
		Name:          "Bob",
		Male:          true,
		Scores:        []int32{2,4,356408,67},
	},nil
}

func main() {
	protocol.RegisterCodec(new(ProtoBufCodec))
	server := server.NewServer(server.WithAddressServer(":1234"))
	err := server.Elem(new(Hello))
	if err != nil {
		panic(err)
	}
	err = server.Start()
	if err != nil {
		panic(err)
	}
	client1,err := client.NewClient(client.WithAddressClient(":1234"),
		client.WithClientCodec("protobuf"),client.WithClientEncoder("text"))
	if err != nil {
		panic(err)
	}
	student := &Student{
		Name:          "Tony",
		Male:          true,
		Scores:        []int32{20,10,20},
	}
	p1 := NewHelloProxy(client1)
	s,err := p1.SayHelloToProtoBuf(student)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	client2,err := client.NewClient(client.WithAddressClient(":1234"))
	if err != nil {
		panic(err)
	}
	student.Name = "Jeni"
	p2 := NewHelloProxy(client2)
	s, err = p2.SayHelloToJson(student)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}