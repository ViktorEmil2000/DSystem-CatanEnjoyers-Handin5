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

	result, err := startBidding(Client, userId, highestBid, MoneyAmount)
	if err != nil {
		log.Fatalf("Could not get result %s:", err)
	}

	if result.ID == userId {
		fmt.Printf("Yes my bid is the higest! my ID was: %s", userId)
	} else {
		fmt.Printf("My was not high enough! my ID was: %s", userId)
	}
}

func startBidding(Client auctionBidder.CommunicationClient, userId int64, highestBid int64, MoneyAmount int64) (*auctionBidder.Result, error) {
	for {
		time.Sleep(100)
		result, err := Client.Result(context.Background(), &auctionBidder.Empty{})
		if err != nil {
			log.Fatalf("Could not get receive %s:", err)
		}
		if result.AuctionOver {
			return result, err
		}

		if result.AuctionActive {
			if userId != result.ID {
				highestBid = result.Amount
				log.Printf("Received message from server: %s", result.Comment)

				if highestBid > MoneyAmount {
					log.Printf("%s: with %s I can't afford to bid higher than the current bid: %s, so I'm outta here", userId, MoneyAmount, highestBid)
					return result, err
				} else {
					Client.Bid(context.Background(), &auctionBidder.FromBidder{
						Amount:    sendBid(highestBid),
						ID:        userId,
						Timestamp: time.Now().Unix(),
					})
				}

			} else {
				fmt.Print("I am currently the highest bidder, and do not need to bid")
				time.Sleep(100)
			}
		}
	}
}

func sendBid(highestBid int64) int64 {
	var bidAmount = highestBid + int64(rand.Intn(900)+100)
	return bidAmount
}

func queryAuctionServer() {

}
