package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/ecode"
	pb "kratos-article/api"
	"kratos-article/internal/biz"
	"kratos-article/internal/dao"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.ArticleServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
	biz biz.Biz
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHelloURL bm demo func.
func (s *Service) ArticleDetail(ctx context.Context, req *pb.ArticleReq) (reply *pb.ArticleResp, err error) {
	articleContentSource, err := s.biz.ArticleDetail(ctx, req.Id)

	if err != nil {
		if errors.Is(err, dao.NotFound) {
			ec := ecode.Int(-404)
			return &pb.ArticleResp{}, ec
		}

	}
	reply = &pb.ArticleResp{
		ArticleId:      articleContentSource.ArticleId,
		ArticleTitle:   articleContentSource.ArticleTitle,
		ArticleAuthor:  articleContentSource.ArticleAuthor,
		ArticleContent: articleContentSource.ArticleContent,
	}
	return
}

//查询文章列表
//func(s *Service) ArticlePage(ctx context.Context, req *pb.ArticlePageReq) (resp *pb.ArticlePageResp, err error){
//
//}
// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
