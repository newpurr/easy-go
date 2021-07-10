package service

import (
	"context"
	"github.com/newpurr/easy-go/application"

	otgorm "github.com/eddycjy/opentracing-gorm"

	"github.com/newpurr/easy-go/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, application.DBEngine))
	return svc
}
