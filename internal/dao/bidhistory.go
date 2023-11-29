package dao

import "github.com/HarryLuo227/simple-bidding-system/internal/model"

func (d *Dao) GetBidHistoryListByAID(id uint32) ([]*model.History, error) {
	bidHistory := model.History{
		AuctionID: id,
	}
	return bidHistory.GetListByAID(d.engine)
}
