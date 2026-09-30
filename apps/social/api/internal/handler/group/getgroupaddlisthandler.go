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

// 获取群申请列表
func Get_group_add_listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupAddListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewGet_group_add_listLogic(r.Context(), svcCtx)
		resp, err := l.Get_group_add_list(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
