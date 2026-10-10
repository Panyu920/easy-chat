package configserver

import (
	"errors"
	"github.com/zeromicro/go-zero/core/conf"
)

var ErrNotSetConfig = errors.New("config server not set")

type ChangeCallback func(bytes []byte) error
type IConfigServer interface {
	Build() error
	SetChangeCallback(callback ChangeCallback)
	GetBytesFromJson() ([]byte, error)
}

type ConfigServer struct {
	IConfigServer
	filePath string
}

func NewConfigServer(filePath string, s IConfigServer) *ConfigServer {
	return &ConfigServer{
		filePath:      filePath,
		IConfigServer: s,
	}
}

func (s *ConfigServer) MustLoad(v any, onchange ChangeCallback) error {

	if s.filePath == "" && s.IConfigServer == nil {
		return ErrNotSetConfig
	}

	if s.IConfigServer == nil {
		//使用go-zero的默认
		conf.MustLoad(s.filePath, v)
		return nil
	}
	//设置配置更新时执行的方法
	if onchange != nil {
		s.IConfigServer.SetChangeCallback(onchange)
	}
	//构建配置更新
	if err := s.IConfigServer.Build(); err != nil {
		return err
	}

	data, err := s.IConfigServer.GetBytesFromJson()
	if err != nil {
		return err
	}
	return LoadFromJsonBytes(data, v)
}

func LoadFromJsonBytes(data []byte, v any) error {
	return conf.LoadFromJsonBytes(data, v)
}
