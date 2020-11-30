package cookie_session

import "sync"

//	memSession 内存对应的 Session 对象
type memSession struct {
	//	全局唯一标识的 session id 对象
	id string
	//	session 数据
	data map[string]interface{}
	//	session 过期时间
	expired int
	//	读写锁，支持多线程
	rwLock sync.RWMutex
}

func NewMemSession(id string) *memSession {
	return &memSession{
		id:   id,
		data: make(map[string]interface{}, 8),
	}
}
