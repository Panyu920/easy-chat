package logic

import (
	"context"

	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/im/rpc/internal/svc"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatLogLogic {
	return &GetChatLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话记录
func (l *GetChatLogLogic) GetChatLog(in *im.GetChatLogReq) (*im.GetChatLogResp, error) {
	// todo: add your logic here and delete this line
	// 1.根据消息ID查询
	if in.MsgId != "" {
		msg, err := l.svcCtx.ChatLogModel.FindOne(l.ctx, in.MsgId)
		if err != nil {
			// 消息不存在
			if err == immodel.ErrNotFound {
				l.Logger.Infof("msgId: %s not found", in.MsgId)
				return &im.GetChatLogResp{
					List: []*im.ChatLog{},
				}, nil
			}
			return nil, errors.Wrapf(xerr.NewDBError(), "find chat log failed %v, msgId: %s", err, in.MsgId)
		} else {
			return &im.GetChatLogResp{
				List: []*im.ChatLog{
					convertDBChatLogToRpcChatLog(msg),
				},
			}, nil
		}
	}
	// 2.根据消息起止时间查询
	msgs, err := l.svcCtx.ChatLogModel.ListChatLogBySendTime(l.ctx, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	if err != nil {
		if err == immodel.ErrNotFound {
			return &im.GetChatLogResp{
				List: []*im.ChatLog{},
			}, nil
		}
		return nil, errors.Wrapf(xerr.NewDBError(), "list chat log failed %v, conversationId: %s, startTime: %d, endTime: %d, limit: %d", err, in.ConversationId, in.StartSendTime, in.EndSendTime, in.Count)
	}

	// 转换为RPC模型
	var rpcMsgs []*im.ChatLog = make([]*im.ChatLog, 0, len(msgs))

	for _, msg := range msgs {
		rpcMsgs = append(rpcMsgs, convertDBChatLogToRpcChatLog(msg))
	}

	return &im.GetChatLogResp{
		List: rpcMsgs,
	}, nil
}

func convertDBChatLogToRpcChatLog(msg *immodel.ChatLog) *im.ChatLog {
	return &im.ChatLog{
		Id:             msg.ID.Hex(),
		ConversationId: msg.ConversationId,
		ChatType:       int32(msg.ChatType),
		SendId:         msg.SendId,
		RecvId:         msg.RecvId,
		MsgContent:     msg.MsgContent,
		MsgType:        int32(msg.MsgType),
		SendTime:       msg.SendTime,
	}
}
