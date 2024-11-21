package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ViktorEmil2000/DSystem-CatanEnjoyers-Handin5/auctionBidder"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

type message struct {
	bidamount int64
	userId    int64
	TimeStamp int64
}
type result struct {
	AuctionActive bool
	comment       string
	ID            int64
}

func main() {
	fmt.Println("Enter port:")
	var MoneyAmount = int64(100000)
	var userId = int64(rand.Intn(10000000))
	var highestBid = int64(0)
	go initializeBidder(userId, highestBid, MoneyAmount)
}

func initializeBidder(userId int64, highestBid int64, MoneyAmount int64) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial AuctionServer %s:", err)
	}
	Client := auctionBidder.NewCommunicationClient(conn)

	result, err := Client.Result(context.Background(), &auctionBidder.Empty{})
	if err != nil {
		log.Fatalf("Could not get result %s:", err)
	}
	highestBid = res

	Client.Bid(context.Background(), &auctionBidder.FromBidder{
		Amount:    sendBid(highestBid),
		ID:        userId,
		Timestamp: time.Now().Unix(),
	})
	auctionBidder.FromAuction()

}

func sendBid(highestBid int64) int64 {
	var bidAmount = highestBid + int64(rand.Intn(900)+100)
	return bidAmount

}

func queryAuctionServer() {

}
