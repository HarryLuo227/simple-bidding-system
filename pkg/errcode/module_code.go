package errcode

var (
	ErrorGetAuctionListFail = NewError(20010001, "取得特定競拍資料失敗")
	ErrorCreateAuctionFail  = NewError(20010002, "建立競拍失敗")
	ErrorUpdateAuctionFail  = NewError(20010003, "更新競拍失敗")
	ErrorDeleteAuctionFail  = NewError(20010004, "刪除競拍失敗")

	ErrorGetLastBidHistoryFail = NewError(20020001, "取得競拍歷史下標失敗")
)
