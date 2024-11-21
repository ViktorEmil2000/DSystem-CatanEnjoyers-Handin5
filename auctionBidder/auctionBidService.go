package auctionBidder

import "context"

type AuctionBidderService struct {
}

func (ABS *AuctionBidderService) Bid(context.Context, *FromBidder) (*FromAuction, error) {

}

// Server sends result to all bidders
func (ABS *AuctionBidderService) Result(context.Context, *Empty) (*Result, error) {

}
func (ABS *AuctionBidderService) mustEmbedUnimplementedCommunicationServer() {
	panic("unimplemented")
}
