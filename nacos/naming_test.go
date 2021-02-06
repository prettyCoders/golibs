package nacos

import (
	"github.com/prettyCoders/golibs/utils/testutil"
	"testing"
)

const naming_instance_name = "nacos-lib"

func TestRegisterInstance(t *testing.T) {
	TestInit(t)
	namingInstance := NewDefaultNamingInstance(naming_instance_name, 8080, nil)
	success, err := RegisterInstance(namingInstance)
	testutil.AssertTrue(t, success)
	testutil.AssertNil(t, err)
}

func TestCall(t *testing.T) {
	TestRegisterInstance(t)
	template := NewDefaultHttpTemplate(naming_instance_name, "/", nil)
	resp, err := Call(template)
	testutil.AssertEqual(t, resp, "")
	testutil.AssertNotNil(t, err)
}
