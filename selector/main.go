package main

import (
	"fmt"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry/consul"
	// "github.com/micro/go-micro/client/selector"
)

func main() {
	// type Selector interface {
	// 	Init(opts ...Option) error
	// 	Options() Options
	// 	// Select returns a function which should return the next node
	// 	Select(service string, opts ...SelectOption) (Next, error)
	// 	// Mark sets the success/error against a node
	// 	Mark(service string, node *registry.Node, err error)
	// 	// Reset returns state back to zero for a service
	// 	Reset(service string)
	// 	// Close renders the selector unusable
	// 	Close() error
	// 	// Name of the selector
	// 	String() string
	// }

	cr := consul.NewRegistry()
	s := selector.NewSelector(selector.Registry(cr))
	fmt.Println(s.String())

	n, err := s.Select("a-Test")

	if err != nil {
		println(err)
		return
	}

	count := map[string]int{}
	for i := 0; i < 1000; i++ {
		node, err := n()
		if err != nil {
			println(err)
			continue
		}
		count[node.Id]++
	}

	for k, v := range count {
		println(k, v)
	}

	s.Reset("a-Test")

}
