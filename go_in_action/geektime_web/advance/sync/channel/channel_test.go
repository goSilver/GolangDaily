package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelReceive(t *testing.T) {
	ch := make(chan string, 1)
	go func() {
		data := <-ch
		fmt.Printf("g1 receive %s\n", data)
	}()

	go func() {
		data := <-ch
		fmt.Printf("g2 receive %s\n", data)
	}()
	ch <- "chensh"
	time.Sleep(3 * time.Second)
}

func TestBroker(t *testing.T) {
	b := &Broker{
		consumers: make([]*Consumer, 0, 10),
	}
	c1 := &Consumer{
		ch: make(chan string, 1),
	}
	c2 := &Consumer{
		ch: make(chan string, 1),
	}
	b.Subscribe(c1)
	b.Subscribe(c2)

	b.Produce("hello")
	fmt.Println(<-c1.ch)
	fmt.Println(<-c2.ch)
}
