package main

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry/consul"
	// "github.com/micro/go-micro/client/selector"
)

func main() {
	// type Broker interface {
	// 	Options() Options
	// 	Address() string
	// 	Connect() error
	// 	Disconnect() error
	// 	Init(...Option) error
	// 	Publish(string, *Message, ...PublishOption) error
	// 	Subscribe(string, Handler, ...SubscribeOption) (Subscriber, error)
	// 	String() string
	// }

	b := broker.NewBroker(
		broker.Registry(consul.NewRegistry()),
		// broker.Addrs("127.0.0.1:10002"),
	)

	fmt.Println(b.String())

	if err := b.Connect(); err != nil {
		fmt.Println("Unexpected connect error:", err)
	}

	fmt.Println(b.Address())

	msg := &broker.Message{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
		Body: []byte(`{"message": "Hello World"}`),
	}

	c := 2
	topic := "a_test_broker"
	var subs []broker.Subscriber
	// done := make(chan bool, c)

	for i := 0; i < c; i++ {
		a := i
		// 实际是向registry注册服务
		sub, err := b.Subscribe(topic, func(p broker.Publication) error {
			// done <- true
			m := p.Message()

			fmt.Println("b.Subscribe i:", a, " m.Body:", string(m.Body))
			if string(m.Body) != string(msg.Body) {
				// fmt.Println("Unexpected msg %s, expected %s", string(m.Body), string(msg.Body))
			}

			return nil
		},
		// broker.Queue不设置时代表广播，否则随机选择一个node推送
		// broker.Queue("shared"),
		)
		if err != nil {
			fmt.Println("Unexpected subscribe error: ", err)
		}
		subs = append(subs, sub)
	}

	time.Sleep(time.Second)
	for i := 0; i < 1; i++ {
		// be.StartTimer()
		if err := b.Publish(topic, msg); err != nil {
			fmt.Println("Unexpected publish error: ", err)
		}
		time.Sleep(time.Second)

		// <-done
		// be.StopTimer()
	}

	for _, sub := range subs {
		sub.Unsubscribe()
	}

	if err := b.Disconnect(); err != nil {
		fmt.Println("Unexpected disconnect error: ", err)
	}

}
