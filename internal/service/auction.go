package service

import "github.com/HarryLuo227/simple-bidding-system/internal/model"

type GetAuctionRequest struct {
	ID uint32 `form:"id"`
}
type UpdateAuctionRequest struct {
	ID             uint32 `form:"id"`
	ItemID         uint32 `form:"item_id"`
	InitBidPrice   int    `form:"init_bid_price"`
	LatestBidPrice int    `form:"latest_bid_price"`
	HammerPrice    int    `form:"hammer_price"`
}

func (svc *Service) GetAuction(param *GetAuctionRequest) (*model.Auction, error) {
	return svc.dao.GetAuction(param.ID)
}

// 下標
func (svc *Service) UpdateAuction(param *UpdateAuctionRequest) error {
	return svc.dao.UpdateAuction(param.ID, param.ItemID, param.InitBidPrice, param.LatestBidPrice)
}
