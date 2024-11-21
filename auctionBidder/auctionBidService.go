package auctionBidder

import "context"

var (
	auctionLive bool
	highestbid  bid
	bidders     []bidder
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
	if auctionLive == false {
		auctionLive = true
	}
	check := false
	for _, element := range bidders {
		if (element.id == FromBidder{}.ID) {
			check = true
			element.bids = append(element.bids, bid{FromBidder{}.Amount, FromBidder{}.Timestamp})
		}
	}
	if check == false {
		temp := []bid{}
		temp = append(temp, bid{FromBidder{}.Amount, FromBidder{}.Timestamp})
		bidders = append(bidders, bidder{FromBidder{}.ID, temp})
	}

	gotHighestBid := false
	if (highestbid.bidammount < FromBidder{}.Amount || (highestbid.bidammount == FromBidder{}.Amount && highestbid.timestamp > FromBidder{}.Timestamp)) {
		highestbid = bid{FromBidder{}.ID, FromBidder{}.Timestamp}
		gotHighestBid = true
	}
	if gotHighestBid {
		return &FromAuction{Acknowledgment: gotHighestBid, Comment: "You got the higest bid"}, nil
	} else {
		return &FromAuction{Acknowledgment: gotHighestBid, Comment: "bidNotHighEnough"}, nil
	}

}

// Server sends result to all bidders
func (ABS *AuctionBidderService) Result(context.Context, *Empty) (*Result, error) {

}
func (ABS *AuctionBidderService) mustEmbedUnimplementedCommunicationServer() {
	panic("unimplemented")
}
