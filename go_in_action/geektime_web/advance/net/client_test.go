package net

import (
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		conn.Close()
		return
	}
	respBs := make([]byte, 16)
	_, err = conn.Read(respBs)
	if err != nil {
		conn.Close()
		return
	}
}
