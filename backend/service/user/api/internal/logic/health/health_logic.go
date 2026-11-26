// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package health

import (
	"context"

	"backend/service/user/api/internal/svc"
	"backend/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health() (resp *types.HealthResponse, err error) {
	return &types.HealthResponse{
		Service: l.svcCtx.Config.Name,
		Status:  "ok",
	}, nil
}
