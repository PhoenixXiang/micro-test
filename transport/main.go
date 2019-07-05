package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/micro/go-micro/transport"
	"github.com/micro/go-micro/transport/grpc"
)

func main() {
	// t := transport.NewTransport() // httpTransport
	t := grpc.NewTransport() // grpcTransport

	// fmt.Println(t.String())

	l, err := t.Listen("127.0.0.1:10001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fn := func(sock transport.Socket) {
		defer sock.Close()

		for {
			var m transport.Message
			if err := sock.Recv(&m); err != nil {
				return
			}
			fmt.Println("server: ",string(m.Body))

			if err := sock.Send(&m); err != nil {
				return
			}
		}
	}

	done := make(chan bool)

	go func() {
		if err := l.Accept(fn); err != nil {
			select {
			case <-done:
			default:
				fmt.Println("Unexpected accept err: ", err)
			}
		}
	}()

	wg := new(sync.WaitGroup)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()

			c, err := t.Dial(l.Addr())
			if err != nil {
				fmt.Println("Unexpected dial err: ", err)
			}
			defer c.Close()

			m := transport.Message{
				Header: map[string]string{
					"Content-Type": "application/json",
				},
				Body: []byte(`{"message": "Hello World` + strconv.Itoa(x) + ` "}`),
			}

			if err := c.Send(&m); err != nil {
				fmt.Println("Unexpected send err: ", err)
			}

			var rm transport.Message

			if err := c.Recv(&rm); err != nil {
				fmt.Println("Unexpected recv err:", err)
			}

			fmt.Println(string(m.Body), string(rm.Body))
			// fmt.Println(c.Local(), c.Remote())
		}(i)
	}

	wg.Wait()
	close(done)
}
