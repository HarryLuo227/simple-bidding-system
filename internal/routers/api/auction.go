package api

import (
	"log"

	"github.com/HarryLuo227/simple-bidding-system/global"
	"github.com/HarryLuo227/simple-bidding-system/internal/service"
	"github.com/HarryLuo227/simple-bidding-system/pkg/app"
	"github.com/HarryLuo227/simple-bidding-system/pkg/convert"
	"github.com/gin-gonic/gin"
)

type Auction struct{}

func NewAuction() Auction {
	return Auction{}
}

func (a Auction) Get(c *gin.Context) {
	param := service.GetAuctionRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		log.Printf("app.BindAndValid in GetAuction err: %v", errs)
		return
	}

	svc := service.New(c.Request.Context())
	result, err := svc.GetAuction(&param)
	if err != nil {
		log.Printf("svc.UpdateAuction err: %v", err)
		return
	}

	response.ToResponse(result)
	return
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
	mutexname := "bid-distributed-lock"
	mutex := global.RedisSync.NewMutex(mutexname)
	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		log.Println("Distributed lock Has already locked by some one")
		log.Println(&param)
		panic(err)
	}
	log.Println("Locked")
	log.Println(&param)

	if err := svc.UpdateAuction(&param); err != nil {
		log.Printf("svc.UpdateAuction err: %v", err)
		return
	}

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
	log.Println("Unlocked")

	response.ToResponse("Succeeded")
	return
}
