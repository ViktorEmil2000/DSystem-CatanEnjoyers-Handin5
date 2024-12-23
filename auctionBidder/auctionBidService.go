package auctionBidder

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
)

var (
	auctionLive      bool
	highestbid       bid
	highestbidder    int64
	bidders          []bidder
	Client           CommunicationClient
	AuctionStartTime int
	highestBidLock   sync.Mutex
	biddersLock      sync.Mutex
)
var kill = false
var auctionStateLock sync.Mutex

type bidder struct {
	id   int64
	bids []bid
}

type bid struct {
	bidamount int64
	timestamp int64
}

type AuctionBidderService struct {
	IsLeader bool
}

func (ABS *AuctionBidderService) Initializer(Auctiongivenstart int) {
	AuctionStartTime = Auctiongivenstart
	if ABS.IsLeader {
		conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to dial AuctionServer %s:", err)

		}
		Client = NewCommunicationClient(conn)
	}
	for !auctionLive {
		if AuctionStartTime <= int(time.Now().Unix()) {
			auctionLive = true
		}
	}
}

func (ABS *AuctionBidderService) Bid(ctx context.Context, FromBidder *FromBidder) (*FromAuction, error) {
	if !checkAuctionOver() {
		biddersLock.Lock()
		BidderExists := false
		for _, element := range bidders {
			if element.id == FromBidder.ID {
				BidderExists = true
				element.bids = append(element.bids, bid{FromBidder.Amount, FromBidder.Timestamp})
				break
			}
		}
		if BidderExists == false {
			temp := []bid{}
			temp = append(temp, bid{FromBidder.Amount, FromBidder.Timestamp})
			bidders = append(bidders, bidder{FromBidder.ID, temp})
		}
		biddersLock.Unlock()
		highestBidLock.Lock()
		gotHighestBid := false
		if highestbid.bidamount < FromBidder.Amount || (highestbid.bidamount == FromBidder.Amount && highestbid.timestamp > FromBidder.Timestamp) {
			highestbid = bid{FromBidder.Amount, FromBidder.Timestamp}
			highestbidder = FromBidder.ID
			gotHighestBid = true
		}
		highestBidLock.Unlock()
		if ABS.IsLeader {
			_, err := Client.Bid(ctx, FromBidder)
			if err != nil {
				log.Fatalf("Unable to contact replication server, data package lost: %s", err)
			}
			log.Printf("We sent a datapackage to Replication server :%v , %v", FromBidder.ID, FromBidder.Amount)
		}
		if gotHighestBid {
			log.Printf("we got a new higest bid from bidder with id: %v with the amount %v$", highestbidder, highestbid.bidamount)
			return &FromAuction{Acknowledgment: gotHighestBid, Comment: "You got the highest bid"}, nil
		} else {
			log.Printf("A bid came thru from %v but it wasen't high enough. Current highest :%v$ the bid was on :%v$", FromBidder.ID, highestbid.bidamount, FromBidder.Amount)
			return &FromAuction{Acknowledgment: gotHighestBid, Comment: "bidNotHighEnough"}, nil
		}
	} else {
		return &FromAuction{Acknowledgment: false, Comment: "Auction is over"}, nil
	}

}

func (ABS *AuctionBidderService) Result(context.Context, *Empty) (*Result, error) {
	highestBidLock.Lock()
    defer highestBidLock.Unlock()
	if !auctionLive {
		fmt.Println()
		log.Print("A user queried Result and was told : Auction hasen't started yet!")
		fmt.Println()
		fmt.Println("............................................................................")
		return &Result{AuctionActive: auctionLive, Comment: "Auction hasen't started yet!", AuctionOver: false}, nil
	} else if checkAuctionOver() {
		fmt.Println()
		log.Print("A user queried Result and was told : Auction is over")
		fmt.Println()
		fmt.Println("............................................................................")
		return &Result{AuctionActive: false, Comment: "Auction is over", AuctionOver: true, ID: highestbidder, Amount: highestbid.bidamount}, nil
	} else {
		fmt.Println()
		log.Print("A user queried Result and was told : Auctions is going and we have a new highest bid")
		fmt.Println()
		fmt.Println("............................................................................")
		return &Result{AuctionActive: auctionLive, Comment: "Auctions is going and we have a new highest bid", ID: highestbidder, Amount: highestbid.bidamount, AuctionOver: false}, nil

	}
}

func checkAuctionOver() bool {
	auctionStateLock.Lock()
    defer auctionStateLock.Unlock()
	if int(time.Now().Unix())-AuctionStartTime >= 100 {
		auctionLive = false
		return true
	} else {
		return false
	}
}
func (ABS *AuctionBidderService) mustEmbedUnimplementedCommunicationServer() {
	panic("unimplemented")
}
