// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package health

import (
	"net/http"

	"backend/service/user/api/internal/logic/health"
	"backend/service/user/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := health.NewHealthLogic(r.Context(), svcCtx)
		resp, err := l.Health()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
