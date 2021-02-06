package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
)

type (
	//ConfigInterface 如果想使用nacos的配置中心，配置结构体需要实现此接口
	ConfigInterface interface {
		//重载接口
		Reload(data string) error
		//检查接口
		AlreadyLoaded() bool
	}

	//ConfigInstance 配置实例，必须实现 ConfigInterface
	ConfigInstance struct {
		DataID string
		Group  string
		ConfigInterface
	}

	//ListenHandler 处理配置监听过程中出现的错误
	ListenHandler func(namespace, group, dataId, data string, err error)
)

//PullConfigYaml 拉取yaml格式配置
func PullConfigYaml(instance *ConfigInstance) error {
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: instance.DataID,
		Group:  instance.Group,
	})
	if err != nil {
		panic(err)
	}
	return yaml.Unmarshal([]byte(content), instance.ConfigInterface)
}

//ListenConfig 监听配置更改
func ListenConfig(instance *ConfigInstance, reload bool, listenHandler ListenHandler) error {
	return configClient.ListenConfig(vo.ConfigParam{
		DataId: instance.DataID,
		Group:  instance.Group,
		OnChange: func(namespace, group, dataId, data string) {
			//是否重载变量
			shouldReload := reload
			//因为开启监听也会调用OnChange函数，所以加此判断
			if !instance.ConfigInterface.AlreadyLoaded() {
				shouldReload = false //如果model 没有被加载过,则不需要做重载
			}
			//重载
			if shouldReload { //重载关键代码
				err := instance.ConfigInterface.Reload(data)
				if err != nil {
					listenHandler(namespace, group, dataId, data, err)
					return
				}
			}
		},
	})
}
