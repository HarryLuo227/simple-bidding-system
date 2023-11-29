package api

import (
	"log"

	"github.com/HarryLuo227/simple-bidding-system/internal/service"
	"github.com/HarryLuo227/simple-bidding-system/pkg/app"
	"github.com/HarryLuo227/simple-bidding-system/pkg/convert"
	"github.com/gin-gonic/gin"
)

type Auction struct{}

func NewAuction() Auction {
	return Auction{}
}

func (a Auction) Update(c *gin.Context) {
	param := service.UpdateAuctionRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		log.Printf("app.BindAndValid in UpdateAuction err: %v", errs)
		return
	}

	svc := service.New(c.Request.Context())
	if err := svc.UpdateAuction(&param); err != nil {
		log.Printf("svc.UpdateAuction err: %v", err)
		return
	}

	response.ToResponse("Succeeded")
	return
}
