package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/prettyCoders/golibs/utils"
	"strconv"
)

type (
	//服务实例
	NamingInstance struct {
		Name      string            //服务名
		Port      uint64            //服务端口
		Weight    float64           //服务权重
		Enable    bool              //是否可用
		Healthy   bool              //是否健康
		Ephemeral bool              //是否是临时实例
		Metadata  map[string]string //元数据
	}

	//服务间HTTP调用模板
	HttpTemplate struct {
		Protocol string            //通信协议
		Name     string            //服务名
		Header   map[string]string //header信息
		Method   string            //请求方法
		URI      string            //uri
		ReqBody  interface{}       //请求体
	}
)

//NewDefaultNamingInstance 创建新的默认服务实例
func NewDefaultNamingInstance(name string, port uint64, metadata map[string]string) *NamingInstance {
	return &NamingInstance{
		Name:      name,
		Port:      port,
		Weight:    10,
		Enable:    true,
		Healthy:   true,
		Ephemeral: true,
		Metadata:  metadata,
	}
}

//NewDefaultHttpTemplate 创建新的默认模版
func NewDefaultHttpTemplate(name string, uri string, reqBody interface{}) *HttpTemplate {
	header := make(map[string]string)
	return &HttpTemplate{
		Protocol: "http",
		Name:     name,
		Header:   header,
		Method:   "POST",
		URI:      uri,
		ReqBody:  reqBody,
	}
}

//RegisterInstance 注册实例
func RegisterInstance(instance *NamingInstance) (bool, error) {
	localIp, err := utils.GetLocalIP()
	if err != nil {
		return false, err
	}
	return namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          localIp.String(),
		Port:        instance.Port,
		ServiceName: instance.Name,
		Weight:      instance.Weight,
		Enable:      instance.Enable,
		Healthy:     instance.Healthy,
		Ephemeral:   instance.Ephemeral,
		Metadata:    instance.Metadata,
	})
}

//Call 服务调用
func Call(template *HttpTemplate) (string, error) {
	instance, err := namingClient.SelectOneHealthyInstance(
		vo.SelectOneHealthInstanceParam{ServiceName: template.Name},
	)
	if err != nil {
		return "", err
	}
	url := template.Protocol + "://" + instance.Ip + ":" + strconv.FormatUint(instance.Port, 10) + template.URI
	return utils.Launch(
		&utils.Request{
			Method:      template.Method,
			Header:      template.Header,
			Url:         url,
			RequestBody: template.ReqBody,
		},
	)
}
