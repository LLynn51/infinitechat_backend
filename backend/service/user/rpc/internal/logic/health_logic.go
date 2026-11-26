package logic

import (
	"context"

	"backend/service/user/rpc/internal/svc"
	"backend/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HealthLogic) Health(in *user.HealthRequest) (*user.HealthResponse, error) {
	return &user.HealthResponse{
		Service: l.svcCtx.Config.Name,
		Status:  "ok",
	}, nil
}
