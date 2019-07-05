package main

import (
	"context"
	"fmt"

	proto "github.com/PhoenixXiang/micro-test/example/pb"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	x := make([]int32, 1)
	x[0] = 1
	// x :=[]*proto.Person{&proto.Person{},&proto.Person{},&proto.Person{}}
	// y := make([]*proto.Person, 3)
	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{
		Pers: x,
	})

	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)

}

