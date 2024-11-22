package main

import (
	"log"
	"net"
	"os"

	"github.com/ViktorEmil2000/DSystem-CatanEnjoyers-Handin5/auctionBidder"
	"google.golang.org/grpc"
)

func main() {

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "50051"
	}

	listen, _ := net.Listen("tcp", ":"+Port)
	log.Println("Listening @ : " + Port)

	grpcserver := grpc.NewServer()

	/*
		SETUP THE AUCTION INSTEAD OF BOOTSERVERSERVICE{}
		Done :D
	*/

	cs := auctionBidder.AuctionBidderService{}
	auctionBidder.RegisterCommunicationServer(grpcserver, &cs)

	grpcserver.Serve(listen)

}
