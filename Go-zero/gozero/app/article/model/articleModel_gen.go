// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleFieldNames          = builder.RawFieldNames(&Article{})
	articleRows                = strings.Join(articleFieldNames, ",")
	articleRowsExpectAutoSet   = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	articleRowsWithPlaceHolder = strings.Join(stringx.Remove(articleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMyDbArticleIdPrefix = "cache:myDb:article:id:"
)

type (
	articleModel interface {
		Insert(ctx context.Context, data *Article) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Article, error)
		Update(ctx context.Context, data *Article) error
		Delete(ctx context.Context, id int64) error
	}

	defaultArticleModel struct {
		sqlc.CachedConn
		table string
	}

	Article struct {
		Id      int64          `db:"id"`
		Title   string `db:"title"`
		Content string `db:"content"`
	}
)

func newArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultArticleModel {
	return &defaultArticleModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`article`",
	}
}

func (m *defaultArticleModel) Delete(ctx context.Context, id int64) error {
	myDbArticleIdKey := fmt.Sprintf("%s%v", cacheMyDbArticleIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, myDbArticleIdKey)
	return err
}

func (m *defaultArticleModel) FindOne(ctx context.Context, id int64) (*Article, error) {
	myDbArticleIdKey := fmt.Sprintf("%s%v", cacheMyDbArticleIdPrefix, id)
	var resp Article
	err := m.QueryRowCtx(ctx, &resp, myDbArticleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultArticleModel) Insert(ctx context.Context, data *Article) (sql.Result, error) {
	myDbArticleIdKey := fmt.Sprintf("%s%v", cacheMyDbArticleIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, articleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.Content)
	}, myDbArticleIdKey)
	return ret, err
}

func (m *defaultArticleModel) Update(ctx context.Context, data *Article) error {
	myDbArticleIdKey := fmt.Sprintf("%s%v", cacheMyDbArticleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.Content, data.Id)
	}, myDbArticleIdKey)
	return err
}

func (m *defaultArticleModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheMyDbArticleIdPrefix, primary)
}

func (m *defaultArticleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultArticleModel) tableName() string {
	return m.table
}