package cookie_session

import (
	"fmt"
	"sync"

	uuid "github.com/satori/go.uuid"
)

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

func (m *memSession) ID() string {
	return m.id
}

func (m *memSession) Load() (err error) {
	return
}

func (m *memSession) Get(key string) (value interface{}, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = fmt.Errorf("Invalid key: ")
		return
	}
	return
}

func (m *memSession) Set(key string, value interface{}) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	m.data[key] = value
}

func (m *memSession) Del(key string) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.data, key)
}

func (m *memSession) SetExpired(expired int) {
	m.expired = expired
}

func (m *memSession) Save() {
	panic("implement me")
}

//	MemSessionMgr 内存 Session 管理器
type MemSessionMgr struct {
	session map[string]Session
	rwLock  sync.RWMutex
}

//	NewMemSessionMgr MemSessionMgr 类构造函数
func NewMemSessionMgr() *MemSessionMgr {
	return &MemSessionMgr{
		session: make(map[string]Session, 1024),
	}
}

func (m *MemSessionMgr) Init(addr string, options ...string) (err error) {
	return
}

//	GetSession: get the session by session id
func (m *MemSessionMgr) GetSession(sessionID string) (sd Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	sd, ok := m.session[sessionID]
	if !ok {
		err = fmt.Errorf("Invalid session id: ")
		return
	}
	return
}

func (m *MemSessionMgr) CreateSession() (sd Session) {
	sessionID := uuid.NewV4().String()
	sd = NewMemSession(sessionID)
	m.session[sd.ID()] = sd
	return
}

func (m *MemSessionMgr) Clear(sessionID string) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	delete(m.session, sessionID)
}
