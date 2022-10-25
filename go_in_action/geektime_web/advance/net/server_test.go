package net

import (
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			handle(conn)
		}()
	}
}

func handle(conn net.Conn) {
	reqBs := make([]byte, 8)
	_, err := conn.Read(reqBs)
	if err != nil {
		conn.Close()
		return
	}
	_, err = conn.Write([]byte("hello world"))
	if err != nil {
		conn.Close()
		return
	}
}
