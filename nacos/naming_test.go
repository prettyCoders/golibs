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
	c := CustomLocation{IP: "127.0.0.1", Port: 8082}
	template := NewDefaultHttpTemplate(naming_instance_name, "/", nil, &c)
	resp := ""
	err := Call(resp, template)
	testutil.AssertEqual(t, resp, "")
	testutil.AssertNotNil(t, err)

	template = NewDefaultHttpTemplate(naming_instance_name, "/", nil, nil)
	err = Call(resp, template)
	testutil.AssertEqual(t, resp, "")
	testutil.AssertNotNil(t, err)
}
