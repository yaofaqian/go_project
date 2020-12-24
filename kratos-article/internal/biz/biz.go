package biz

import (
	"kratos-article/internal/dao"
	"kratos-article/internal/model"
	"context"
)

type Biz interface {
	ArticleDetail(ctx context.Context, id int64) (*model.ArticleContent, error)
}
type biz struct {
	dao dao.Dao
}

//查询一篇文章
func (b *biz) ArticleDetail(ctx context.Context, id int64) (res *ArticleContent, err error) {
	articleContentDao, err := b.dao.Article(ctx, id)
	if err != nil {
		return nil, err
	}
	res = &ArticleContent{
		articleContentDao.ArticleId,
		articleContentDao.ArticleTitle,
		articleContentDao.ArticleAuthor,
		articleContentDao.ArticleHits,
		articleContentDao.ArticleContent,
		"张无忌",
		articleContentDao.ArticleUpdateTime,
	}
	return
}
//文章内容
type ArticleContent struct {
	ArticleId         int64
	ArticleTitle      string
	ArticleAuthor     string
	ArticleHits       int64
	ArticleContent    string
	UserName          string
	ArticleUpdateTime string
}
