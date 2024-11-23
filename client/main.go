package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/ViktorEmil2000/DSystem-CatanEnjoyers-Handin5/auctionBidder"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go initializeBidder()
	wg.Wait()
}

func initializeBidder() {
	var MoneyAmount = int64(rand.Intn(5000) + 5000)
	var userId = int64(rand.Intn(10000000))

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial AuctionServer %s:", err)
	}
	Client := auctionBidder.NewCommunicationClient(conn)

	result, err := startBidding(Client, userId, MoneyAmount)
	if err != nil {
		log.Fatalf("Could not get result %s:", err)
	}

	if result.ID == userId {
		log.Printf("Yes my bid is the higest! my ID was: %s", userId)
	} else {
		log.Printf("My was not high enough! my ID was: %s", userId)
	}
}

func startBidding(Client auctionBidder.CommunicationClient, userId int64, MoneyAmount int64) (*auctionBidder.Result, error) {
	for {
		time.Sleep(time.Millisecond * 1000)
		result, err := Client.Result(context.Background(), &auctionBidder.Empty{})
		if err != nil {
			log.Fatalf("Could not get receive %s:", err)
		}
		highestBid := result.Amount
		if result.AuctionOver {
			return result, err
		}

		if result.AuctionActive {
			if userId != result.ID {
				log.Printf("Received message from server: %s", result.Comment)

				if highestBid > MoneyAmount {
					log.Printf("%v: with %v I can't afford to bid higher than the current bid: %v, so I'm outta here", userId, MoneyAmount, highestBid)
					return result, err
				} else {
					var fromAuction, err = sendBid(Client, userId, highestBid)
					if err != nil {
						log.Fatalf("Something went wrong with sendBid-method %s:", err)
					}

					if fromAuction.Acknowledgment {
						log.Printf("Acknowledgement came back positive for: %v: '%v'", userId, fromAuction.Comment)
					} else {
						log.Printf("Acknowledgement came back negative for: %v: '%v'", userId, fromAuction.Comment)
					}
				}
			} else {
				log.Print("I am currently the highest bidder, and do not need to bid")
				time.Sleep(time.Millisecond * 1000)
			}

		} else {
			log.Print("Currently waiting for Auction to start")
		}
	}
}

func sendBid(Client auctionBidder.CommunicationClient, userId int64, highestBid int64) (*auctionBidder.FromAuction, error) {
	var bidAmount = highestBid + int64(rand.Intn(900)+100)
	result, err := Client.Bid(context.Background(), &auctionBidder.FromBidder{
		Amount:    bidAmount,
		ID:        userId,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		log.Fatalf("Could not place a bid %s:", err)
	}

	return result, err
}
