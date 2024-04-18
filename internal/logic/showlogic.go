package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"shortener/internal/svc"
	"shortener/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowReqeust) (resp *types.ShowResponse, err error) {
	// 1. 根据短链接查询原始长链接
	u, err := l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{String: req.ShortUrl, Valid: true})
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.New("短链接不存在")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
