package cookie_session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type SessionMgrType string

const (
	//	SessionId在cookie里面的名字
	SessionCookieName = "session_id"
	//	Session对象在Context里面的名字
	SessionContextName                = "session"
	Memory             SessionMgrType = "memory"
	Redis              SessionMgrType = "redis"
)

//	Session 接口
type Session interface {
	//	获取Session对象的ID
	ID() string
	//	加载 redis 数据到 session data
	Load() error
	//	获取 key 对应的 value 值
	Get(string) (interface{}, error)
	//	设置 key 对应的 value 值
	Set(string, interface{})
	//	删除 key 对应的 value 值
	Del(string)
	//	落盘数据到 redis
	Save()
	//	设置 Redis 数据过期时间，内存版本无效
	SetExpired(int)
}

//	SessionMgr	Session 管理器对象
type SessionMgr interface {
	//	初始化 Redis 数据库连接
	Init(addr string, options ...string) error
	//	通过 SessionID 获取已经初始化的 Session 对象
	GetSession(string) (Session, error)
	//	创建一个新的 Session 对象
	CreateSession() Session
	//	使用 SessionID 清空一个 Session 对象
	Clear(string)
}

//	Options Cookie 对应的相关选项
type Options struct {
	Path   string
	Domain string
	// Cookie中的SessionID存活时间
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HttpOnly bool
}

//	生成 Session 管理器对象
func CreateSessionMgr(name SessionMgrType, addr string, options ...string) (sm SessionMgr, err error) {
	switch name {
	case Memory:
		sm = NewMemSessionMgr()
	case Redis:
		sm = NewRedisSessionMgr()
	default:
		err = fmt.Errorf("unsupported %v\n", name)
		return
	}
	err = sm.Init(addr, options...)
	return
}

func SessionMiddleware(sm SessionMgr, options Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		var session Session
		//	尝试从 cookie 获取 session ID
		sessionID, err := c.Cookie(SessionCookieName)
		if err != nil {
			log.Printf("get session_id from cookie failed, err: %v\n", err)
			session = sm.CreateSession()
			sessionID = session.ID()
		} else {
			log.Printf("SessionId: %v\n", sessionID)
			session, err = sm.GetSession(sessionID)
			if err != nil {
				log.Printf("Get seesion by %s failed, err: %v\n", sessionID, err)
				//session = sm.CreateSession()
				sessionID = session.ID()
			}
		}
		session.SetExpired(options.MaxAge)
		c.Set(SessionContextName, session)
		c.Set("options", options)
		c.Set("sessionID", sessionID)
		c.Set("session", session)
		//defer sm.Clear(sessionID)
		//c.SetCookie(SessionContextName, sessionID, options.MaxAge, options.Path, options.Domain, options.Secure, options.HttpOnly)

		c.Next()
		session.Save()
	}
}
