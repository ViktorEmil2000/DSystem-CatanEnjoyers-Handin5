package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/ViktorEmil2000/DSystem-CatanEnjoyers-Handin5/auctionBidder"
	"google.golang.org/grpc"
)

var (
	port        = "50051"
	userId      int64
	MoneyAmount int64
)

var backupRun = true

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go initializebidderVariables()
	wg.Wait()
}

func initializebidderVariables() {
	MoneyAmount = int64(rand.Intn(10000) + 5000)
	userId = int64(rand.Intn(10000000))
	log.Print("*****************************************************")
	log.Printf("New bidder has been registered with ID: %v", userId)
	log.Print("****************************************************")
	initializeBidder()
	ch := make(chan bool)
	ch <- true

}

func initializeBidder() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		if backupRun {
			log.Printf("Failed to dial AuctionServer %s:", err)
		} else {
			log.Fatalf("Failed to dial AuctionServer %s:", err)
		}
		tryOnBackup()
	}
	Client := auctionBidder.NewCommunicationClient(conn)

	result, err := startBidding(Client, userId, MoneyAmount)
	if err != nil {
		if backupRun {
			log.Printf("Could not get result %s:", err)
		} else {
			log.Fatalf("Could not get result %s:", err)
		}
		tryOnBackup()
	}

	if result.ID == userId {
		log.Printf("%v: Yes I won with the highest bid of %v$", userId, result.Amount)
	} else {
		log.Printf("%v: I didn't win the auction", userId)
	}
}

func startBidding(Client auctionBidder.CommunicationClient, userId int64, MoneyAmount int64) (*auctionBidder.Result, error) {
	for {
		fmt.Println("............................................................")
		time.Sleep(time.Millisecond * 1000)
		result, err := Client.Result(context.Background(), &auctionBidder.Empty{})
		if err != nil {
			if backupRun {
				log.Printf("Could not get receive %s:", err)
			} else {
				log.Fatalf("Could not get receive %s:", err)
			}
			tryOnBackup()
		}
		if result.AuctionOver {
			return result, err
		}

		if result.AuctionActive {
			if userId != result.ID {
				highestBid := result.Amount
				log.Printf("Received message from server: %s: %v$", result.Comment, highestBid)

				if highestBid+1000 > MoneyAmount {
					log.Printf("%v: with only %v$ my spouse won't forgive me if I use more money than the highest bid: %v$, so I'm outta here", userId, MoneyAmount, highestBid)
					return result, err
				} else {
					var fromAuction, err = sendBid(Client, userId, highestBid)
					if err != nil {
						if backupRun {
							log.Printf("Something went wrong with sendBid-method %s:", err)
						} else {
							log.Fatalf("Something went wrong with sendBid-method %s:", err)
						}
						tryOnBackup()
					}

					if fromAuction.Acknowledgment {
						fmt.Println()
						fmt.Println("*********************************************************")
						fmt.Printf("Acknowledgement came back positive for: %v: '%v'", userId, fromAuction.Comment)
						fmt.Println()
						fmt.Println("*********************************************************")
						fmt.Println()
					} else {
						fmt.Println()
						fmt.Println("*********************************************************")
						fmt.Printf("***Acknowledgement came back negative for: %v: '%v'***", userId, fromAuction.Comment)
						fmt.Println()
						fmt.Println("*********************************************************")
						fmt.Println()
					}
				}
			} else {
				log.Print("I am currently the highest bidder, and do not need to bid")
				time.Sleep(time.Millisecond * 3000)
			}

		} else {
			log.Print("Currently waiting for Auction to start")
		}
	}
}

func tryOnBackup() {
	port = "50052"
	log.Print("We lost connection on my server and is trying the backup")
	backupRun = false
	initializeBidder()
}

func sendBid(Client auctionBidder.CommunicationClient, userId int64, highestBid int64) (*auctionBidder.FromAuction, error) {
	var bidAmount = highestBid + int64(rand.Intn(900)+100)
	log.Printf("Hmmm I can afford to bid higher than %v$... I WANT TO BID %v$!!!", highestBid, bidAmount)
	result, err := Client.Bid(context.Background(), &auctionBidder.FromBidder{
		Amount:    bidAmount,
		ID:        userId,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		log.Printf("Could not place a bid %s:", err)
	}

	return result, err
}
