package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Course struct {
	arr []int
}

func (c *Course) Hello(_ *struct{}, res *string) error {
	fmt.Println("Call to Hello")
	*res = "Hello!"
	return nil
}

func (c *Course) Add(v int, _ *struct{}) error {
	fmt.Println("Call to Add")
	time.Sleep(time.Second)
	c.arr = append(c.arr, v)
	return nil
}

func (c *Course) GetArray(_ *struct{}, res *[]int) error {
	fmt.Println("Call to GetArray")
	time.Sleep(10 * time.Second)
	*res = c.arr
	return nil
}

func (c *Course) Reset(_ *struct{}, _ *struct{}) error {
	fmt.Println("Call to Reset")
	c.arr = []int{}
	return nil
}

func main() {
	add, err := net.ResolveTCPAddr("tcp", ":8888")
	fmt.Println("Litening on", add)
	CheckErr(err)

	listener, err := net.ListenTCP("tcp", add)
	CheckErr(err)

	course := new(Course)
	rpc.Register(course)
	rpc.Accept(listener)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
