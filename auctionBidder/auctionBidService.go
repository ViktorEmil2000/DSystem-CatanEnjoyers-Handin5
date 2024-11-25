package auctionBidder

import (
	"context"
	"fmt"
	"log"
	"time"
)

var (
	auctionLive   bool
	highestbid    bid
	highestbidder int64
	bidders       []bidder
)
var AuctionStartTime = 0

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

func (ABS *AuctionBidderService) Bid(ctx context.Context, FromBidder *FromBidder) (*FromAuction, error) {

	if !checkAuctionOver() {
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

		gotHighestBid := false
		if highestbid.bidamount < FromBidder.Amount || (highestbid.bidamount == FromBidder.Amount && highestbid.timestamp > FromBidder.Timestamp) {
			highestbid = bid{FromBidder.Amount, FromBidder.Timestamp}
			highestbidder = FromBidder.ID
			gotHighestBid = true
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
	if !auctionLive {
		AuctionStartTime++
		if AuctionStartTime >= 5 {
			auctionLive = true
			AuctionStartTime = int(time.Now().Unix())
		}
	}

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
	if int(time.Now().Unix())-AuctionStartTime >= 50 {
		auctionLive = false
		return true
	} else {
		return false
	}
}
func (ABS *AuctionBidderService) mustEmbedUnimplementedCommunicationServer() {
	panic("unimplemented")
}
