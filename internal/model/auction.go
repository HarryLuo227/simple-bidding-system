package model

import (
	"github.com/HarryLuo227/simple-bidding-system/global"
	"github.com/jinzhu/gorm"
)

type Auction struct {
	*Model
	ItemID         uint32 `json:"item_id"`
	InitBidPrice   int    `json:"init_bid_price"`
	LatestBidPrice int    `json:"latest_bid_price"`
	HammerPrice    int    `json:"hammer_price"`
}

func (a Auction) TableName() string {
	return global.DatabaseSetting.TablePrefix + "auction"
}

// Query
func (a Auction) Get(db *gorm.DB) (*Auction, error) {
	err := db.Where("id = ? AND is_del = ?", a.ID, 0).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// Update
func (a Auction) Update(db *gorm.DB, values map[string]int) error {
	err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}
