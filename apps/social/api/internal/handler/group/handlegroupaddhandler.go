// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package group

import (
	"net/http"

	"easy-chat/apps/social/api/internal/logic/group"
	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 处理群申请
func Handle_group_addHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupAddHandlerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewHandle_group_addLogic(r.Context(), svcCtx)
		resp, err := l.Handle_group_add(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
