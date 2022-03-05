package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Get struct {
	ClientId string
	OpId     uint32
	Key      string
}

type Response struct {
	OpId  uint32
	GId   uint64
	Key   string
	Value string
}

type TailServer struct {
	kv map[string]string
}

func NewTailServer() *TailServer {
	return &TailServer{
		kv: make(map[string]string),
	}
}

func (t *TailServer) GetValue(req Get, res *Response) error {
	time.Sleep(3 * time.Second)
	fmt.Println("Call to GetValue")
	value, exist := t.kv[req.Key]
	if !exist {
		return errors.New("Key does not exist")
	}
	*res = Response{req.OpId, 1234, req.Key, value}
	fmt.Println(res)
	return nil
}

func main() {
	add, err := net.ResolveTCPAddr("tcp", ":9990")
	fmt.Println("Litening on", add)
	CheckErr(err)

	listener, err := net.ListenTCP("tcp", add)
	CheckErr(err)

	tailServer := NewTailServer()
	tailServer.kv["key1"] = "value1"
	tailServer.kv["key2"] = "value2"
	tailServer.kv["key3"] = "value3"
	err = rpc.Register(tailServer)
	CheckErr(err)
	rpc.Accept(listener)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
