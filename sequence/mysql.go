package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 建立mysql连接，执行replace info语句
// REPLACE INFO sequence(stub) VALUES('a');
// SELECT LAST_INSERT_ID();
const (
	sqlReplaceInfoStub = `replace INTO sequence (stub) VALUES ('a');`
)

type MySQL struct {
	conn sqlx.SqlConn
}

func NewMySQL(dsn string) *MySQL {
	return &MySQL{
		conn: sqlx.NewMysql(dsn),
	}
}

// Next 生成下一个id
func (m *MySQL) Next() (uint64, error) {
	stmt, err := m.conn.Prepare(sqlReplaceInfoStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	defer stmt.Close()
	rest, err := stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	id, err := rest.LastInsertId()
	if err != nil {
		logx.Errorw("rest.LastInsertId failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	return uint64(id), nil
}
