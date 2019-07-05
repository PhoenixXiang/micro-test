package main

import (
	"fmt"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	// type Registry interface {
	// 	Init(...Option) error
	// 	Options() Options
	// 	Register(*Service, ...RegisterOption) error
	// 	Deregister(*Service) error
	// 	GetService(string) ([]*Service, error)
	// 	ListServices() ([]*Service, error)
	// 	Watch(...WatchOption) (Watcher, error)
	// 	String() string
	// }
	cr := consul.NewRegistry()
	// fmt.Println(cr.String())

	s := &registry.Service{
		Name:    "a-Test",
		Version: "1.0.0",
		Metadata:map[string]string{
			"a":"test",
		},
		Nodes: []*registry.Node{
			{
				Id:      "test-node",
				Address: "node-address",
				Port:    10001,
			},
		},
	}
	e := cr.Register(s)
	if e != nil {
		fmt.Println(e)
	}

	// e := cr.Deregister(s)
	// if e != nil {
	// 	fmt.Println(e)
	// }

	// ss, err := cr.ListServices()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// ss, _ := cr.GetService("a-Test")
	// for _, s := range ss {
	// 	fmt.Println(s.Name)
	// 	fmt.Println(s.Version)
	// 	fmt.Println(s.Metadata)
	// 	for _, e := range s.Endpoints {
	// 		fmt.Println(e)
	// 	}
	// 	for _, n := range s.Nodes {
	// 		fmt.Println(n)
	// 	}
	// }

}
