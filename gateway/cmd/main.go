package main

import (
	"common/api"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load(".env")

	// Create Grpc Connection
	conn := GetGrpcConn(":2000")

	defer conn.Close()

	// Add client to GRPC server
	c := api.NewOrderServiceClient(conn)

	// Port
	port := os.Getenv("PORT")

	if port == ""{
	    log.Fatal("PORT not found")
	}

	http := NewHttpHandler(c)

	r := http.mount()

	log.Printf("Server running on port %v", port)

	log.Fatal(http.run(":" + port, r))
}