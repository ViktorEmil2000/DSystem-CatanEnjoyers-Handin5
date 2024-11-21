package auctionBidder

import "context"

type auctionBidderService struct {
}

func (ABS *auctionBidderService) Bid(context.Context, *FromBidder) (*FromAuction, error) {

}

// Server sends result to all bidders
func (ABS *auctionBidderService) Result(context.Context, *Empty) (*Result, error) {

}
func (ABS *auctionBidderService) mustEmbedUnimplementedCommunicationServer() {
	panic("unimplemented")
}
