package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"shortener/internal/svc"
	"shortener/internal/types"
	"shortener/pkg/connect"
	"shortener/pkg/extract"
	"shortener/pkg/urlTool"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Convert 将长链接转换为短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// todo: add your logic here and delete this line
	// 1. 校验输入参数
	//	1.1 数据不能为空
	//	1.2 输入的长链接能够正常访问
	if ok := connect.Get(req.LongUrl); !ok {
		return nil, errors.New("无效的链接")
	}
	// 	1.3 长链接是否已经转换过
	md5Value := extract.Sum([]byte(req.LongUrl))
	u, err := l.svcCtx.ShortUrlModel.FindOneByMd5(l.ctx, sql.NullString{
		String: md5Value,
		Valid:  true,
	})
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, fmt.Errorf("该链接已经转换过，短链接为：%s", u.Surl.String)
		}
		logx.Errorw("ShortUrlModel.FindOneByMd5 failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	basePath, err := urlTool.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urlTool.GetBasePath failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	_, err = l.svcCtx.ShortUrlModel.FindOneBySurl(l.ctx, sql.NullString{
		String: basePath,
		Valid:  true,
	})
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, errors.New("该链接已经是短链")
		}
		logx.Errorw("ShortUrlModel.FindOneBySurl failed", logx.LogField{Key: "err", Value: err.Error()})
		return nil, err
	}
	// 2. 生成短链接
	// 3. 保存短链接
	// 4. 返回响应
	return
}
