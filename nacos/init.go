package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	namingClient naming_client.INamingClient
	configClient config_client.IConfigClient
	Model        NacosModel
)

type NacosModel string

const (
	DEBUG   NacosModel = "debug"
	RELEASE NacosModel = "release"
)

//Init 初始化nacos客户端配置
func Init(model NacosModel, sc []constant.ServerConfig, cc *constant.ClientConfig) error {
	Model = model
	//初始化服务发现客户端
	if c, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	); err != nil {
		return err
	} else {
		namingClient = c
	}
	//初始化配置中心客户端
	if c, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	); err != nil {
		return err
	} else {
		configClient = c
	}
	return nil
}
