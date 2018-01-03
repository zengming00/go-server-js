package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

type Session struct {
	mSessionID        string
	mLastTimeAccessed time.Time
	mValues           map[string]interface{}
}

type SessionMgr struct {
	mCookieName     string       // 客户端cookie名称
	mLock           sync.RWMutex // 互斥(保证线程安全)
	mMaxLifeTimeSec int64        // 垃圾回收时间(秒)

	mSessions map[string]*Session
}

func NewSessionMgr(cookieName string, maxLifeTimeSec int64) *SessionMgr {
	mgr := &SessionMgr{mCookieName: cookieName, mMaxLifeTimeSec: maxLifeTimeSec, mSessions: make(map[string]*Session)}
	go mgr.GC()
	return mgr
}

func (mgr *SessionMgr) StartSession(w http.ResponseWriter, r *http.Request) string {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	cookie, err := r.Cookie(mgr.mCookieName)
	if err == nil {
		mgr.mLock.Lock()
		defer mgr.mLock.Unlock()

		sessionID := cookie.Value
		if session, ok := mgr.mSessions[sessionID]; ok {
			session.mLastTimeAccessed = time.Now()
			return sessionID
		}
	}
	newSessionID := url.QueryEscape(mgr.NewSessionID())
	session := &Session{mSessionID: newSessionID, mLastTimeAccessed: time.Now(), mValues: make(map[string]interface{})}
	mgr.mSessions[newSessionID] = session
	cookie = &http.Cookie{Name: mgr.mCookieName, Value: newSessionID, Path: "/", HttpOnly: true, MaxAge: int(mgr.mMaxLifeTimeSec)}
	http.SetCookie(w, cookie)
	return newSessionID
}

func (mgr *SessionMgr) EndSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(mgr.mCookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	delete(mgr.mSessions, cookie.Value)

	// 让浏览器cookie立刻过期
	cookie = &http.Cookie{Name: mgr.mCookieName, Path: "/", HttpOnly: true, MaxAge: -1}
	http.SetCookie(w, cookie)
}

func (mgr *SessionMgr) EndSessionByID(sessionID string) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	delete(mgr.mSessions, sessionID)
}

func (mgr *SessionMgr) SetSessionVal(sessionID string, key string, value interface{}) {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	if session, ok := mgr.mSessions[sessionID]; ok {
		session.mValues[key] = value
	}
}

func (mgr *SessionMgr) GetSessionVal(sessionID string, key string) (value interface{}, ok bool) {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()

	if session, ok := mgr.mSessions[sessionID]; ok {
		value, ok = session.mValues[key]
	}
	return
}

func (mgr *SessionMgr) GetSessionIDList() []string {
	mgr.mLock.RLock()
	defer mgr.mLock.RUnlock()

	sessionIDList := make([]string, len(mgr.mSessions))
	i := 0
	for k := range mgr.mSessions {
		sessionIDList[i] = k
		i++
	}
	return sessionIDList
}

func (mgr *SessionMgr) GC() {
	mgr.mLock.Lock()
	defer mgr.mLock.Unlock()

	for sessionID, session := range mgr.mSessions {
		if session.mLastTimeAccessed.Unix()+mgr.mMaxLifeTimeSec < time.Now().Unix() {
			delete(mgr.mSessions, sessionID)
		}
	}
	time.AfterFunc(time.Duration(mgr.mMaxLifeTimeSec)*time.Second, func() { mgr.GC() })
}

func (mgr *SessionMgr) NewSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		nano := time.Now().UnixNano()
		return strconv.FormatInt(nano, 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}
