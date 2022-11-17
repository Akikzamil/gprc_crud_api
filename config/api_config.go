package config

import (
	"fmt"
	"grpc_practice/crud_proto"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

type server struct {
	crud_proto.UnsafeUserServiceServer
}

func APIConfig() {
	fmt.Println("Api Service Started")
	lis,err := net.Listen("tcp","0.0.0.0:10000")
	if err != nil {
		log.Fatalf("Failed to listen: %v",err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	crud_proto.RegisterUserServiceServer(s,&server{})
	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis);err != nil {
			log.Fatalf("failed to serve: %v",err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch,os.Interrupt)

	<-ch

	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stoping the listener")
	lis.Close()
	fmt.Println("End of Program")
}