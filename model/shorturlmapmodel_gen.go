// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shortUrlMapFieldNames          = builder.RawFieldNames(&ShortUrlMap{})
	shortUrlMapRows                = strings.Join(shortUrlMapFieldNames, ",")
	shortUrlMapRowsExpectAutoSet   = strings.Join(stringx.Remove(shortUrlMapFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	shortUrlMapRowsWithPlaceHolder = strings.Join(stringx.Remove(shortUrlMapFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	shortUrlMapModel interface {
		Insert(ctx context.Context, data *ShortUrlMap) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ShortUrlMap, error)
		FindOneByMd5(ctx context.Context, md5 sql.NullString) (*ShortUrlMap, error)
		FindOneBySurl(ctx context.Context, surl sql.NullString) (*ShortUrlMap, error)
		Update(ctx context.Context, data *ShortUrlMap) error
		Delete(ctx context.Context, id int64) error
	}

	defaultShortUrlMapModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ShortUrlMap struct {
		Id       int64          `db:"id"`        // 主键
		CreateAt time.Time      `db:"create_at"` // 创建时间
		CreateBy string         `db:"create_by"` // 创建者
		IsDel    int64          `db:"is_del"`    // 是否删除：0正常1删除
		Lurl     sql.NullString `db:"lurl"`      // 长链接
		Md5      sql.NullString `db:"md5"`       // ⻓链接MD5
		Surl     sql.NullString `db:"surl"`      // 短链接
	}
)

func newShortUrlMapModel(conn sqlx.SqlConn) *defaultShortUrlMapModel {
	return &defaultShortUrlMapModel{
		conn:  conn,
		table: "`short_url_map`",
	}
}

func (m *defaultShortUrlMapModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultShortUrlMapModel) FindOne(ctx context.Context, id int64) (*ShortUrlMap, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shortUrlMapRows, m.table)
	var resp ShortUrlMap
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlMapModel) FindOneByMd5(ctx context.Context, md5 sql.NullString) (*ShortUrlMap, error) {
	var resp ShortUrlMap
	query := fmt.Sprintf("select %s from %s where `md5` = ? limit 1", shortUrlMapRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, md5)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlMapModel) FindOneBySurl(ctx context.Context, surl sql.NullString) (*ShortUrlMap, error) {
	var resp ShortUrlMap
	query := fmt.Sprintf("select %s from %s where `surl` = ? limit 1", shortUrlMapRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, surl)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShortUrlMapModel) Insert(ctx context.Context, data *ShortUrlMap) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, shortUrlMapRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CreateBy, data.IsDel, data.Lurl, data.Md5, data.Surl)
	return ret, err
}

func (m *defaultShortUrlMapModel) Update(ctx context.Context, newData *ShortUrlMap) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shortUrlMapRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.CreateBy, newData.IsDel, newData.Lurl, newData.Md5, newData.Surl, newData.Id)
	return err
}

func (m *defaultShortUrlMapModel) tableName() string {
	return m.table
}
