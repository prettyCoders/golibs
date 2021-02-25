package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/prettyCoders/golibs/utils/testutil"
	"testing"
)

func TestInit(t *testing.T) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := *constant.NewClientConfig(
		constant.WithTimeoutMs(500),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("nacos-dir/log/"),
		constant.WithCacheDir("nacos-dir/cache/"),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
	)
	testutil.AssertNil(t, Init(RELEASE, sc, &cc))
}
