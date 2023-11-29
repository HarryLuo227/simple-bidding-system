package model

import "github.com/jinzhu/gorm"

type History struct {
	*Model
	AuctionID uint32 `json:"auction_id"`
	BidPrice  int    `json:"bid_price"`
}

func (h History) GetLastBidHistory(db *gorm.DB) error {
	err := db.Where("auction_id = ? AND is_del = ?", h.AuctionID, 0).Last(&h).Error
	if err != nil {
		return err
	}
	return nil
}
