package model

import (
	"github.com/HarryLuo227/simple-bidding-system/global"
	"github.com/jinzhu/gorm"
)

type History struct {
	*Model
	AuctionID uint32 `json:"auction_id"`
	BidPrice  int    `json:"bid_price"`
}

func (h History) TableName() string {
	return global.DatabaseSetting.TablePrefix + "bidhistory"
}

func (h History) GetLastBidHistory(db *gorm.DB) (*History, error) {
	err := db.Where("auction_id = ? AND is_del = ?", h.AuctionID, 0).Last(&h).Error
	if err != nil {
		return nil, err
	}
	return &h, nil
}

// Insert
func (h History) Create(db *gorm.DB) error {
	err := db.Create(&h).Error
	if err != nil {
		return err
	}
	return nil
}
