package dao

import "github.com/HarryLuo227/simple-bidding-system/internal/model"

func (d *Dao) GetAuction(id uint32) (*model.Auction, error) {
	auction := model.Auction{
		Model: &model.Model{
			ID: id,
		},
	}
	return &auction, auction.Get(d.engine)
}

func (d *Dao) UpdateAuction(id, itemId uint32, initBidPrice, latestBidPrice int) error {
	auction := model.Auction{
		Model: &model.Model{
			ID: id,
		},
	}

	values := make(map[string]int)
	if itemId > 0 {
		values["item_id"] = int(itemId)
	}
	if initBidPrice >= 0 {
		values["init_bid_price"] = initBidPrice
	}
	// Get Last BidPrice
	history := model.History{
		AuctionID: id,
	}
	if err := history.GetLastBidHistory(d.engine); err != nil {
		return err
	}
	if latestBidPrice > history.BidPrice {
		values["latest_bid_price"] = latestBidPrice
	}

	// Any update operation need to use transaction to ensure data racing.
	tx := d.engine.Begin()
	if err := auction.Update(tx, values); err != nil {
		tx.Rollback()
		return err
	} else {
		tx.Commit()
		return nil
	}
}
