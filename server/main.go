package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/ViktorEmil2000/DSystem-CatanEnjoyers-Handin5/auctionBidder"
	"google.golang.org/grpc"
)

func main() {

	AuctionStartTime := int(time.Now().Unix()) + 1
	go staging(true, AuctionStartTime)
	go staging(false, AuctionStartTime)

	ch := make(chan bool)
	ch <- true
}

func staging(isleader bool, AuctionStartTime int) {
	if isleader {
		go makeServer(isleader, AuctionStartTime)
		for int(time.Now().Unix())-AuctionStartTime <= 40 {

		}
		log.Print("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::..")
		log.Print("The main Server is now dead!")
		log.Print("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::....")
	} else {
		go makeServer(isleader, AuctionStartTime)
		ch := make(chan bool)
		ch <- true
	}
}

func makeServer(isLeader bool, AuctionStartTime int) {
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
	go cs.Initializer(AuctionStartTime)
	auctionBidder.RegisterCommunicationServer(grpcserver, &cs)

	grpcserver.Serve(listen)

}
