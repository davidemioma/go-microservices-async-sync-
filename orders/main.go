package main

func main(){
	grpcServer := NewRPCServer(":2000")

	grpcServer.Run()
}