package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Manager struct{
		cooikeName	string
		lock		sync.Mutex
		provider	Provider
		maxLifeTime int64
}


type Provider interface {
	SessionInit(sid string) (Session, error)  //实现Session的初始化，操作成功则返回此新的Session变量
	SessionRead(sid string) (Session, error)  //读取sessionID; 若不存在，则创建
	SessionDestory(sid string) error          //销毁session
	SessionGC(macLifeTime int64) 			//删除过期数据
}

type Session interface {
	Set(key, val interface{}) error     //由key设置session
	Get(key interface{}) interface{}    //由指定key获取session Id
	Delete(key interface{}) error
	SessionID() string         			//返回当前session Id
}


var provides = make(map[string]Provider)

func NewManger(provideName, cooikename string, maxLiefTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("sessoin: 未知的session提供者 %q (是否忘记导入？)", provideName)
	}
	return &Manager{provider:provider, cooikeName:cooikename, maxLifeTime:maxLiefTime}, nil
}

func Register(name string, provider Provider)  {
	if provider == nil {       //参数空缺
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {     //重复注册
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

/**
随机生成唯一的session Id
 */
func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

/**
检测是否已经有某个Session与当前来访用户发生了关联，如果没有则创建之。
 */
func (manager *Manager) SessionStart(w http.ResponseWriter, r * http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cooike, err := r.Cookie(manager.cooikeName)
	if err != nil || cooike.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cooike := http.Cookie{	Name: 		manager.cooikeName,
								Value: 		url.QueryEscape(sid),
								Path:		"/",
								HttpOnly:	true,
								MaxAge: 	int(manager.maxLifeTime)}
		http.SetCookie(w, &cooike)
	}else {
		sid, _ := url.QueryUnescape(cooike.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return session
}

/**
删除session的同时，还告诉客户端cookie已经删除
 */
func (manager *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cooikeName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		manager.provider.SessionDestory(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name:			 manager.cooikeName,
								Path:		"/",
								HttpOnly:	true,
								Expires:	expiration,
								MaxAge:		-1}
		http.SetCookie(w, &cookie)
	}
}


func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {manager.GC()})
}