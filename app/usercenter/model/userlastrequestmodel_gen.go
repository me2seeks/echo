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
	userLastRequestFieldNames          = builder.RawFieldNames(&UserLastRequest{})
	userLastRequestRows                = strings.Join(userLastRequestFieldNames, ",")
	userLastRequestRowsExpectAutoSet   = strings.Join(stringx.Remove(userLastRequestFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userLastRequestRowsWithPlaceHolder = strings.Join(stringx.Remove(userLastRequestFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserLastRequestUserIdPrefix = "cache:userLastRequest:userId:"
)

type (
	userLastRequestModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *UserLastRequest) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*UserLastRequest, error)
		Update(ctx context.Context, session sqlx.Session, data *UserLastRequest) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserLastRequest) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		DeleteSoft(ctx context.Context, session sqlx.Session, data *UserLastRequest) error
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*UserLastRequest, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserLastRequest, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserLastRequest, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserLastRequest, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserLastRequest, error)
		Delete(ctx context.Context, session sqlx.Session, userId int64) error
	}

	defaultUserLastRequestModel struct {
		sqlc.CachedConn
		table string
	}

	UserLastRequest struct {
		UserId          int64     `db:"user_id"`           // 用户ID
		LastRequestTime time.Time `db:"last_request_time"` // 用户最后一次请求时间
		CreateAt        time.Time `db:"create_at"`         // 创建时间
		UpdateAt        time.Time `db:"update_at"`         // 更新时间
		DeleteAt        time.Time `db:"delete_at"`         // 删除时间
		DelState        int64     `db:"del_state"`         // 删除状态
		Version         uint64    `db:"version"`           // 版本号
	}
)

func newUserLastRequestModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserLastRequestModel {
	return &defaultUserLastRequestModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_last_request`",
	}
}

func (m *defaultUserLastRequestModel) Delete(ctx context.Context, session sqlx.Session, userId int64) error {
	userLastRequestUserIdKey := fmt.Sprintf("%s%v", cacheUserLastRequestUserIdPrefix, userId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, userId)
		}
		return conn.ExecCtx(ctx, query, userId)
	}, userLastRequestUserIdKey)
	return err
}
func (m *defaultUserLastRequestModel) FindOne(ctx context.Context, userId int64) (*UserLastRequest, error) {
	userLastRequestUserIdKey := fmt.Sprintf("%s%v", cacheUserLastRequestUserIdPrefix, userId)
	var resp UserLastRequest
	err := m.QueryRowCtx(ctx, &resp, userLastRequestUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and del_state = ? limit 1", userLastRequestRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId, globalkey.DelStateNo)
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

func (m *defaultUserLastRequestModel) Insert(ctx context.Context, session sqlx.Session, data *UserLastRequest) (sql.Result, error) {
	data.DeleteAt = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	userLastRequestUserIdKey := fmt.Sprintf("%s%v", cacheUserLastRequestUserIdPrefix, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userLastRequestRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UserId, data.LastRequestTime, data.DeleteAt, data.DelState, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.UserId, data.LastRequestTime, data.DeleteAt, data.DelState, data.Version)
	}, userLastRequestUserIdKey)
}

func (m *defaultUserLastRequestModel) Update(ctx context.Context, session sqlx.Session, data *UserLastRequest) (sql.Result, error) {
	userLastRequestUserIdKey := fmt.Sprintf("%s%v", cacheUserLastRequestUserIdPrefix, data.UserId)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, userLastRequestRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.LastRequestTime, data.DeleteAt, data.DelState, data.Version, data.UserId)
		}
		return conn.ExecCtx(ctx, query, data.LastRequestTime, data.DeleteAt, data.DelState, data.Version, data.UserId)
	}, userLastRequestUserIdKey)
}

func (m *defaultUserLastRequestModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *UserLastRequest) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	userLastRequestUserIdKey := fmt.Sprintf("%s%v", cacheUserLastRequestUserIdPrefix, data.UserId)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ? and version = ? ", m.table, userLastRequestRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.LastRequestTime, data.DeleteAt, data.DelState, data.Version, data.UserId, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.LastRequestTime, data.DeleteAt, data.DelState, data.Version, data.UserId, oldVersion)
	}, userLastRequestUserIdKey)
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

func (m *defaultUserLastRequestModel) DeleteSoft(ctx context.Context, session sqlx.Session, data *UserLastRequest) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteAt = time.Now()
	if err := m.UpdateWithVersion(ctx, session, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "UserLastRequestModel delete err : %+v", err)
	}
	return nil
}

func (m *defaultUserLastRequestModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultUserLastRequestModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultUserLastRequestModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*UserLastRequest, error) {

	builder = builder.Columns(userLastRequestRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserLastRequest
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserLastRequestModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserLastRequest, error) {

	builder = builder.Columns(userLastRequestRows)

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

	var resp []*UserLastRequest
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserLastRequestModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*UserLastRequest, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(userLastRequestRows)

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

	var resp []*UserLastRequest
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultUserLastRequestModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*UserLastRequest, error) {

	builder = builder.Columns(userLastRequestRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserLastRequest
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserLastRequestModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*UserLastRequest, error) {

	builder = builder.Columns(userLastRequestRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserLastRequest
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserLastRequestModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

func (m *defaultUserLastRequestModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
func (m *defaultUserLastRequestModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserLastRequestUserIdPrefix, primary)
}
func (m *defaultUserLastRequestModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and del_state = ? limit 1", userLastRequestRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultUserLastRequestModel) tableName() string {
	return m.table
}
