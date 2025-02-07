package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGrpcConn(addr string) *grpc.ClientConn  {
	conn, err := grpc.NewClient(":2000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to grpc server: %v", err)
	}

	log.Println("Connected to GRPC server at port 2000")

	defer conn.Close()

	return conn
}