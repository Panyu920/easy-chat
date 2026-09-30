// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package friend

import (
	"net/http"

	"easy-chat/apps/social/api/internal/logic/friend"
	"easy-chat/apps/social/api/internal/svc"
	"easy-chat/apps/social/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 处理好友申请（同意/拒绝）
func Handle_friend_addHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendAddHandlerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := friend.NewHandle_friend_addLogic(r.Context(), svcCtx)
		resp, err := l.Handle_friend_add(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
