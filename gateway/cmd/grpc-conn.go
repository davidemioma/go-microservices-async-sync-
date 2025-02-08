package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetGrpcConn(addr string) *grpc.ClientConn  {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect to grpc server: %v", err)
	}

	log.Printf("Connected to GRPC server at port %v", addr)

	return conn
}