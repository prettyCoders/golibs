package request_session

import (
	"github.com/petermattis/goid"
	"sync"
)

var requestSessions = sync.Map{}

// Set 设置一个 Request session
func Set(session interface{}) {

	goID := getGoID()
	requestSessions.Store(goID, session)
}

// Get 返回设置的 Request session
func Get() (interface{}, bool) {
	goID := getGoID()
	return requestSessions.Load(goID)
}

// Delete 删除设置的 Request session
func Delete() {
	goID := getGoID()
	requestSessions.Delete(goID)
}

//getGoID 获取当前协程ID
func getGoID() int64 {
	return goid.Get()
}
