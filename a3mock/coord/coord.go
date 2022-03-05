package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Coord struct {
	headServerIPPort string
	tailServerIPPort string
}

func NewCoord() *Coord {
	return &Coord{
		headServerIPPort: "127.0.0.1:8888",
		tailServerIPPort: "127.0.0.1:9990",
	}
}

func (c *Coord) GetHeadServer(_ *struct{}, res *string) error {
	fmt.Println("Call to GetHeadServer")
	*res = c.headServerIPPort
	return nil
}

func (c *Coord) GetTailServer(_ *struct{}, res *string) error {
	fmt.Println("Call to GetTailServer")
	*res = c.tailServerIPPort
	return nil
}

func main() {
	add, err := net.ResolveTCPAddr("tcp", ":7777")
	fmt.Println("Litening on", add)
	CheckErr(err)

	listener, err := net.ListenTCP("tcp", add)
	CheckErr(err)

	coord := NewCoord()
	err = rpc.Register(coord)
	CheckErr(err)
	rpc.Accept(listener)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
