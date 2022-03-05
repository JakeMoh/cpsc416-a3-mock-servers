package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type RequestPut struct { // RequestPut is used to send put request to head server and is expected to receive from tail server
	ClientId              string
	OpId                  uint32
	Key                   string
	Value                 string
	LocalTailServerIPPort string // TODO Note: The client should wait for a response from tail server and the tail server should send ASK message
}

type Response struct {
	OpId  uint32
	GId   uint64
	Key   string
	Value string
}

type HeadServer struct {
	kv map[string]string
}

func NewHeadServer() *HeadServer {
	return &HeadServer{
		kv: make(map[string]string),
	}
}

func (h *HeadServer) PutKeyValue(req RequestPut, res *Response) error { // TODO
	time.Sleep(3 * time.Second)
	fmt.Println("Call to PutKeyValue")
	h.kv[req.Value] = req.Value
	*res = Response{req.OpId, 1234, req.Key, req.Value}
	client, err := rpc.Dial("tcp", req.LocalTailServerIPPort)
	CheckErr(err)
	err = client.Call("KVS.ReceivePutACK", res, nil)
	fmt.Println("Response:")
	fmt.Println(res)
	CheckErr(err)
	return nil
}

func main() {
	add, err := net.ResolveTCPAddr("tcp", ":8888")
	fmt.Println("Litening on", add)
	CheckErr(err)

	listener, err := net.ListenTCP("tcp", add)
	CheckErr(err)

	headServer := NewHeadServer()
	err = rpc.Register(headServer)
	CheckErr(err)
	rpc.Accept(listener)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
