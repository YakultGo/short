package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"
)

type ServiceContext struct {
	Config        config.Config
	ShortUrlModel model.ShortUrlMapModel
	Sequence      sequence.Sequence
	Blacklist     map[string]struct{}
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	m := make(map[string]struct{})
	for _, v := range c.ShortUrlBlacklist {
		m[v] = struct{}{}
	}
	return &ServiceContext{
		Config:        c,
		ShortUrlModel: model.NewShortUrlMapModel(conn),
		Sequence:      sequence.NewMySQL(c.Sequence.DSN),
		Blacklist:     m,
	}
}
