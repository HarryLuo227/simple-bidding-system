package service

import "github.com/HarryLuo227/simple-bidding-system/internal/model"

type GetBidHistoryListByAIDRequest struct {
	AID uint32 `form:"auction_id"`
}

func (svc *Service) GetBidHistoryListByAID(param *GetBidHistoryListByAIDRequest) ([]*model.History, error) {
	return svc.dao.GetBidHistoryListByAID(param.AID)
}
