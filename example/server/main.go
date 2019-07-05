package main

import (
	"context"
	"fmt"
	"strconv"

	proto "github.com/PhoenixXiang/micro-test/example/pb"
	micro "github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + strconv.Itoa(int(req.Pers[0]))
	fmt.Println(req.Pers)
	// fmt.Println(req.Pers[1])
	// a := req.Pers[1]
	// a.Name = "Test"
	// fmt.Println(req.Pers[1])
	// fmt.Println(len(req.Pers), cap(req.Pers))
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
