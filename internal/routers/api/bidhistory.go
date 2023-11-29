package api

import (
	"log"

	"github.com/HarryLuo227/simple-bidding-system/internal/service"
	"github.com/HarryLuo227/simple-bidding-system/pkg/app"
	"github.com/HarryLuo227/simple-bidding-system/pkg/convert"
	"github.com/gin-gonic/gin"
)

type History struct{}

func NewHistory() History {
	return History{}
}

func (h History) GetBidHistoryListByAID(c *gin.Context) {
	param := service.GetBidHistoryListByAIDRequest{
		AID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		log.Printf("app.BindAndValid in GetBidHistoryListByAID err: %v", errs)
		return
	}

	svc := service.New(c.Request.Context())
	result, err := svc.GetBidHistoryListByAID(&param)
	if err != nil {
		log.Printf("svc.GetBidHistoryListByAID err: %v", err)
		return
	}

	response.ToResponse(result)
	return
}
