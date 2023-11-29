package api

import "github.com/gin-gonic/gin"

type Auction struct{}

func NewAuction() Auction {
	return Auction{}
}

func (a Auction) Update(c *gin.Context) {

}
