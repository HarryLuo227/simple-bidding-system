package routers

import (
	"github.com/HarryLuo227/simple-bidding-system/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Auction
	auction := api.NewAuction()
	r.GET("/bid/:id", auction.Get)
	r.POST("/bid/:id", auction.Update)

	return r
}
