package main

import (
	"fmt"
	"log"
	"net/rpc"
)

var orderCh chan int
var result chan int
var queue chan int

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Add(client *rpc.Client, v int) {
	go func() {
		if len(queue) == cap(queue) {
			fmt.Println("queue capacity is reached")
			return
		}
		queue <- 0
		orderCh <- 0
		err := client.Call("Course.Add", v, nil)
		CheckErr(err)
		result <- v
		<-orderCh
		<-queue
	}()
}

func GetArrayLen(client *rpc.Client) {
	go func() {
		if len(queue) == cap(queue) {
			fmt.Println("queue capacity is reached")
			return
		}
		queue <- 0
		orderCh <- 0
		var arr []int
		err := client.Call("Course.GetArray", struct{}{}, &arr)
		CheckErr(err)
		result <- len(arr)
		<-orderCh
		<-queue
	}()
}

func main() {
	client, err := rpc.Dial("tcp", ":8888")
	CheckErr(err)
	orderCh = make(chan int, 1)
	queue = make(chan int, 11)
	result = make(chan int, 1)

	Add(client, -1)
	Add(client, -1)
	Add(client, -1)
	Add(client, -1)
	Add(client, -1)

	Add(client, -1)
	Add(client, -1)
	Add(client, -1)
	Add(client, -1)
	Add(client, -1)

	GetArrayLen(client)

	//<-orderCh

	//time.Sleep(100 * time.Second)
	for c := range result {
		fmt.Println(c)
	}

	//err = conn.Call("Course.Reset", struct{}{}, nil)
	//CheckErr(err)
	//
	//var reply string
	//err = client.Call("Course.Hello", struct{}{}, &reply)
	//CheckErr(err)
	//
	//
	//
	//err = client.Call("Course.Add", 3, nil)
	//CheckErr(err)
	//
	//err = client.Call("Course.Add", 7, nil)
	//CheckErr(err)
	//
	//var arr []int
	//err = client.Call("Course.GetArray", struct{}{}, &arr)
	//CheckErr(err)

	//fmt.Println(reply)
	//fmt.Println(arr)
}
