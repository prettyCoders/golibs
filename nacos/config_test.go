package nacos

import (
	"fmt"
	"github.com/prettyCoders/golibs/utils/testutil"
	"testing"
)

type MyConfig struct {
	Name string
	Age  uint
}

func (c *MyConfig) Reload(data string) error {
	fmt.Println("reload data", data)
	return nil
}

//检查接口
func (c *MyConfig) AlreadyLoaded() bool {
	return c.Name != "" && c.Age > 0
}

func TestPullConfigYaml(t *testing.T) {
	TestInit(t)
	myConfig := MyConfig{}
	configInstance := ConfigInstance{
		DataID:          "nacos-lib",
		Group:           "nacos-lib",
		ConfigInterface: &myConfig,
	}
	err := PullConfigYaml(&configInstance)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, myConfig.Name, "sunlight")
	testutil.AssertTrue(t, myConfig.Age == 26)
}

func TestListenConfig(t *testing.T) {
	TestInit(t)
	myConfig := MyConfig{}
	configInstance := ConfigInstance{
		DataID:          "nacos-lib",
		Group:           "nacos-lib",
		ConfigInterface: &myConfig,
	}
	err := PullConfigYaml(&configInstance)
	testutil.AssertNil(t, err)
	err = ListenConfig(&configInstance, true, func(namespace, group, dataId, data string, err error) {
		fmt.Println("listen err", err)
	})
	testutil.AssertNil(t, err)
	select {}
}
