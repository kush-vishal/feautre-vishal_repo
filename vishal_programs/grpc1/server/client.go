package main

import (
	"context"
	"fmt"
	"log"

	"grpc1/server/message/hellopb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	c, err := grpc.Dial("localhost:8080", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	client := hellopb.NewHelloServiceClient(c)
	request := &hellopb.HelloRequest{Name: " john"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
}
