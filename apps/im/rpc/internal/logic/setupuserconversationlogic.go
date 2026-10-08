package logic

import (
	"context"

	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/rpc/im"
	"easy-chat/apps/im/rpc/internal/svc"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/wuid"
	"easy-chat/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetUpUserConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 建立会话:  私聊
func (l *SetUpUserConversationLogic) SetUpUserConversation(in *im.SetUpUserConversationReq) (*im.SetUpUserConversationResp, error) {
	// todo: add your logic here and delete this line

	// 1.生成会话ID
	conversationID := wuid.CombineUserID(in.RecvId, in.SendId)
	// 2.查询会话是否存在
	conversation, err := l.svcCtx.ConversationModel.FindByConversationId(l.ctx, conversationID)
	if err != nil && err != immodel.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewDBError(), "ConversationModel.FindByConversationId err %v, conversationID %s", err, conversationID)
	}
	if conversation != nil {
		return &im.SetUpUserConversationResp{Conversation: &im.Conversation{
			ConversationId: conversationID,
			ChatType:       int32(conversation.ChatType),
			UserId:         in.SendId,
		}}, nil
	}

	// 会话不存在, 则创建会话
	conversation = &immodel.Conversation{
		ConversationId: conversationID,
		ChatType:       constant.ChatType(in.ChatType),
	}
	err = l.svcCtx.ConversationModel.WithTransaction(l.ctx, func(ctx context.Context) error {
		// 1. 插入会话
		err := l.svcCtx.ConversationModel.Insert(ctx, conversation)
		if err != nil {
			return errors.Wrapf(xerr.NewDBError(), "ConversationModel.Insert failed err %v, conversationID %s", err, conversationID)
		}
		// 2. 更新或插入会话到用户会话表
		err = l.SetUpUserOneConversation(conversationID, in.SendId, conversation.ChatType, false)
		if err != nil {
			return errors.Wrapf(xerr.NewDBError(), "SetUpUserOneConversation failed err %v, conversationID %s", err, conversationID)
		}
		err = l.SetUpUserOneConversation(conversationID, in.RecvId, conversation.ChatType, false)
		if err != nil {
			return errors.Wrapf(xerr.NewDBError(), "SetUpUserOneConversation failed err %v, conversationID %s", err, conversationID)
		}
		// 3.在rediss设置会话ID和会话seq
		_, err = l.svcCtx.RedisxClient.Setnx(conversationID, "0")
		if err != nil {
			return errors.Wrapf(xerr.NewDBError(), "RedisxClient.Setnx err %v, conversationID %s", err, conversationID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &im.SetUpUserConversationResp{
		Conversation: &im.Conversation{
			ConversationId: conversationID,
			ChatType:       int32(conversation.ChatType),
			UserId:         in.SendId,
		},
	}, nil
}

func (l *SetUpUserConversationLogic) SetUpUserOneConversation(conversationID, userId string, chatType constant.ChatType, isShow bool) error {
	conversations, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, userId)
	if err != nil && err != immodel.ErrNotFound {
		return err
	}
	if conversations != nil {
		// 用户会话存在, 更新会话列表
		conversations.ConversationList[conversationID] = &immodel.Conversation{
			ConversationId: conversationID,
			ChatType:       chatType,
			IsShow:         isShow,
		}
		_, err = l.svcCtx.ConversationsModel.Update(l.ctx, conversations)
		if err != nil {
			return err
		}
	}
	conversations = &immodel.Conversations{}
	// 用户会话不存在, 则 插入会话到用户会话表
	conversations.ConversationList = make(map[string]*immodel.Conversation)
	conversations.ConversationList[conversationID] = &immodel.Conversation{
		ConversationId: conversationID,
		ChatType:       chatType,
		IsShow:         isShow,
	}
	conversations.UserId = userId
	err = l.svcCtx.ConversationsModel.Insert(l.ctx, conversations)
	if err != nil {
		return err
	}
	return nil
}
