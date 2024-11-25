package main

import (
	"log"
	"net"
	"os"

	"github.com/ViktorEmil2000/DSystem-CatanEnjoyers-Handin5/auctionBidder"
	"google.golang.org/grpc"
)

func main() {
	go makeServer(true)
	go makeServer(false)

	ch := make(chan bool)
	ch <- true
}

func makeServer(isLeader bool) {
	Port := os.Getenv("PORT")
	if Port == "" {
		if isLeader {
			Port = "50051"
		} else {
			Port = "50052"
		}
	}

	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {

	}
	log.Println("Listening @ : " + Port)

	grpcserver := grpc.NewServer()

	/*
		SETUP THE AUCTION INSTEAD OF BOOTSERVERSERVICE{}
		Done :D
	*/

	cs := auctionBidder.AuctionBidderService{IsLeader: isLeader}
	auctionBidder.RegisterCommunicationServer(grpcserver, &cs)

	grpcserver.Serve(listen)

}
