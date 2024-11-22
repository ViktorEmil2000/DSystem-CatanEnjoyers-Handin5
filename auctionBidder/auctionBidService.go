package auctionBidder

import "context"

var (
	auctionLive   bool
	highestbid    bid
	highestbidder int64
	bidders       []bidder
)

type bidder struct {
	id   int64
	bids []bid
}

type bid struct {
	bidammount int64
	timestamp  int64
}

type AuctionBidderService struct {
}

func (ABS *AuctionBidderService) Bid(context.Context, *FromBidder) (*FromAuction, error) {
	if !auctionLive {
		auctionLive = true
	}
	BidderExists := false
	for _, element := range bidders {
		if (element.id == FromBidder{}.ID) {
			BidderExists = true
			element.bids = append(element.bids, bid{FromBidder{}.Amount, FromBidder{}.Timestamp})
			break
		}
	}
	if BidderExists == false {
		temp := []bid{}
		temp = append(temp, bid{FromBidder{}.Amount, FromBidder{}.Timestamp})
		bidders = append(bidders, bidder{FromBidder{}.ID, temp})
	}

	gotHighestBid := false
	if (highestbid.bidammount < FromBidder{}.Amount || (highestbid.bidammount == FromBidder{}.Amount && highestbid.timestamp > FromBidder{}.Timestamp)) {
		highestbid = bid{FromBidder{}.Amount, FromBidder{}.Timestamp}
		highestbidder = FromBidder{}.ID
		gotHighestBid = true
	}

	if gotHighestBid {
		return &FromAuction{Acknowledgment: gotHighestBid, Comment: "You got the highest bid"}, nil
	} else {
		return &FromAuction{Acknowledgment: gotHighestBid, Comment: "bidNotHighEnough"}, nil
	}

}

func (ABS *AuctionBidderService) Result(context.Context, *Empty) (*Result, error) {
	if !auctionLive {
		return &Result{AuctionActive: auctionLive, Comment: "Auction hasen't started yet!"}, nil
	} else {
		return &Result{AuctionActive: auctionLive, Comment: "Auctions is going and we have a highest bid", ID: highestbidder, Ammount: highestbid.bidammount}, nil
	}
}

func (ABS *AuctionBidderService) mustEmbedUnimplementedCommunicationServer() {
	panic("unimplemented")
}
