// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"time"

	"github.com/Masterminds/squirrel"
	"github.com/me2seeks/echo-hub/common/globalkey"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysNotificationsFieldNames          = builder.RawFieldNames(&SysNotifications{})
	sysNotificationsRows                = strings.Join(sysNotificationsFieldNames, ",")
	sysNotificationsRowsExpectAutoSet   = strings.Join(stringx.Remove(sysNotificationsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysNotificationsRowsWithPlaceHolder = strings.Join(stringx.Remove(sysNotificationsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSysNotificationsIdPrefix = "cache:sysNotifications:id:"
)

type (
	sysNotificationsModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *SysNotifications) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysNotifications, error)
		Update(ctx context.Context, session sqlx.Session, data *SysNotifications) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *SysNotifications) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *SysNotifications) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*SysNotifications, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*SysNotifications, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*SysNotifications, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*SysNotifications, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*SysNotifications, error)
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultSysNotificationsModel struct {
		sqlc.CachedConn
		table string
	}

	SysNotifications struct {
		Id       int64          `db:"id"`
		Message  sql.NullString `db:"message"`
		CreateAt time.Time      `db:"create_at"`
		UpdateAt time.Time      `db:"update_at"`
		DeleteAt time.Time      `db:"delete_at"`
		DelState int64          `db:"del_state"`
		Version  uint64         `db:"version"` // 版本号
	}
)

func newSysNotificationsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysNotificationsModel {
	return &defaultSysNotificationsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_notifications`",
	}
}

func (m *defaultSysNotificationsModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	sysNotificationsIdKey := fmt.Sprintf("%s%v", cacheSysNotificationsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, sysNotificationsIdKey)
	return err
}
func (m *defaultSysNotificationsModel) FindOne(ctx context.Context, id int64) (*SysNotifications, error) {
	sysNotificationsIdKey := fmt.Sprintf("%s%v", cacheSysNotificationsIdPrefix, id)
	var resp SysNotifications
	err := m.QueryRowCtx(ctx, &resp, sysNotificationsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", sysNotificationsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysNotificationsModel) Insert(ctx context.Context, session sqlx.Session, data *SysNotifications) (sql.Result, error) {
	data.DeleteAt = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	sysNotificationsIdKey := fmt.Sprintf("%s%v", cacheSysNotificationsIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, sysNotificationsRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Message, data.DeleteAt, data.DelState, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.Message, data.DeleteAt, data.DelState, data.Version)
	}, sysNotificationsIdKey)
}

func (m *defaultSysNotificationsModel) Update(ctx context.Context, session sqlx.Session, data *SysNotifications) (sql.Result, error) {
	sysNotificationsIdKey := fmt.Sprintf("%s%v", cacheSysNotificationsIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysNotificationsRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Message, data.DeleteAt, data.DelState, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Message, data.DeleteAt, data.DelState, data.Version, data.Id)
	}, sysNotificationsIdKey)
}

func (m *defaultSysNotificationsModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *SysNotifications) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	sysNotificationsIdKey := fmt.Sprintf("%s%v", cacheSysNotificationsIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, sysNotificationsRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Message, data.DeleteAt, data.DelState, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.Message, data.DeleteAt, data.DelState, data.Version, data.Id, oldVersion)
	}, sysNotificationsIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}

func (m *defaultSysNotificationsModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *SysNotifications) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteAt = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "SysNotificationsModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultSysNotificationsModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultSysNotificationsModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultSysNotificationsModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*SysNotifications, error) {

	builder = builder.Columns(sysNotificationsRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysNotifications
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysNotificationsModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*SysNotifications, error) {

	builder = builder.Columns(sysNotificationsRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysNotifications
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysNotificationsModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*SysNotifications, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(sysNotificationsRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*SysNotifications
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultSysNotificationsModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*SysNotifications, error) {

	builder = builder.Columns(sysNotificationsRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysNotifications
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysNotificationsModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*SysNotifications, error) {

	builder = builder.Columns(sysNotificationsRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*SysNotifications
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysNotificationsModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultSysNotificationsModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultSysNotificationsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSysNotificationsIdPrefix, primary)
}
func (m *defaultSysNotificationsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", sysNotificationsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultSysNotificationsModel) tableName() string {
	return m.table
}