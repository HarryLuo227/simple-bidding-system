package service

type CreateAuctionRequest struct {
	BidID int
}

type UpdateAuctionRequest struct {
	ID        uint32
	LatestBid int
}

func (svc *Service) CreateAuction() {

}

func (svc *Service) UpdateAuction() {

}
