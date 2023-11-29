package service

import (
	"context"

	"github.com/HarryLuo227/simple-bidding-system/global"
	"github.com/HarryLuo227/simple-bidding-system/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) *Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return &svc
}
