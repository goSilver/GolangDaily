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
