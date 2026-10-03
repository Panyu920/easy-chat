package msgtransfer

import (
	"context"
	immodel "easy-chat/apps/im/im_model"
	"easy-chat/apps/im/ws/websocketx"
	"easy-chat/apps/task/mq/internal/svc"
	"easy-chat/apps/task/mq/mqtype"
	"easy-chat/pkg/constant"
	"easy-chat/pkg/wuid"
	"encoding/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgChatTransfer struct {
	svc *svc.ServiceContext
	logx.Logger
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		svc:    svc,
		Logger: logx.WithContext(context.Background()),
	}
}

func (m *MsgChatTransfer) Consume(ctx context.Context, key, value string) error {
	m.Logger.Infof("MsgChatTransferSuccess key :%s , val :%s", key, value)
	var data mqtype.MqChatType
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		m.Logger.Errorf("Unmarshal chat failed: %v", err)
		return err
	}
	// 保存聊天记录
	err = m.SaveChatLog(ctx, &data)
	if err != nil {
		m.Logger.Errorf("Save chat log failed: %v", err)
		return err
	}

	// 将聊天记录发送到im服务
	var frame websocketx.Message
	frame.FrameType = websocketx.FrameTypeData
	frame.Method = "push"
	frame.FromID = constant.REDIS_KEY_USER_ID
	frame.Data = data
	err = m.svc.WsClient.Send(frame)
	if err != nil {
		m.Logger.Errorf("Send frame failed: %v", err)
		return err
	}

	return nil
}

func (m *MsgChatTransfer) SaveChatLog(ctx context.Context, data *mqtype.MqChatType) error {
	if data.ConversationId == "" {
		data.ConversationId = wuid.CombineUserID(data.SendId, data.RecvId)
	}

	err := m.svc.ChatLogModel.Insert(ctx, &immodel.ChatLog{
		ConversationId: data.ConversationId,
		SendId:         data.SendId,
		RecvId:         data.RecvId,
		ChatType:       data.ChatType,
		MsgType:        data.MsgType,
		MsgContent:     data.Content,
		SendTime:       time.Now().Unix(),
		FrameFrom:      data.SendId,
		Status:         0,
	})
	if err != nil {
		m.Logger.Errorf("Insert chat log failed: %v", err)
		return err
	}

	return nil
}
