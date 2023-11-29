package dao

import "github.com/HarryLuo227/simple-bidding-system/internal/model"

func (d *Dao) GetAuction(id uint32) (*model.Auction, error) {
	auction := model.Auction{
		Model: &model.Model{
			ID: id,
		},
	}
	return auction.Get(d.engine)
}

func (d *Dao) UpdateAuction(id, itemId uint32, initBidPrice, latestBidPrice int) error {
	auction := model.Auction{
		Model: &model.Model{
			ID: id,
		},
	}
	history := model.History{
		AuctionID: id,
	}

	values := make(map[string]int)
	if itemId > 0 {
		values["item_id"] = int(itemId)
	}
	if initBidPrice >= 0 {
		values["init_bid_price"] = initBidPrice
	}
	// Get Last BidPrice
	lastBidHistory, err := history.GetLastBidHistory(d.engine)
	if err != nil {
		return err
	}
	if latestBidPrice > lastBidHistory.BidPrice {
		values["latest_bid_price"] = latestBidPrice
	}

	// Any update operation need to use transaction to ensure data racing.
	tx := d.engine.Begin()
	if err := auction.Update(tx, values); err != nil {
		tx.Rollback()
		tx.Commit()
		return err
	} else {
		tx.Commit()
	}

	// If latest_bid_price in auction table updates successfully, then create new bid history.
	if latestBidPrice > lastBidHistory.BidPrice {
		history.BidPrice = latestBidPrice
		if err := history.Create(d.engine); err != nil {
			return err
		}
	}

	return nil
}
