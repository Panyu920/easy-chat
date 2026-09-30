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

// 获取好友申请列表
func Get_friend_add_listHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendAddListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := friend.NewGet_friend_add_listLogic(r.Context(), svcCtx)
		resp, err := l.Get_friend_add_list(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
