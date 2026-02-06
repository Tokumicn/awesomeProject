package service

import (
	"i_intent/data/model"
	"sync"
	"time"
)

var _sessionManager *SessionManager

type Session struct {
	SessionID string                 `json:"session_id"`
	Messages  []models.ChatMessage   `json:"messages"`
	Context   map[string]interface{} `json:"context"`
	CreatedAt *time.Time             `json:"created_at"`
}

type SessionManager struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

func NewSessionManager() *SessionManager {
	if _sessionManager == nil {
		_sessionManager = &SessionManager{
			sessions: make(map[string]*Session),
		}
	}
	return _sessionManager
}

func (sm *SessionManager) GetOrCreate(sessionID string) (string, *Session) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if sessionID == "" {
		sessionID = generateSessionID()
	}

	if _, exists := sm.sessions[sessionID]; !exists {
		sm.sessions[sessionID] = &Session{
			SessionID: sessionID,
			Messages:  []models.ChatMessage{},
			Context:   make(map[string]interface{}),
		}
	}

	return sessionID, sm.sessions[sessionID]
}

func (sm *SessionManager) Reset(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, sessionID)
}

func generateSessionID() string {
	return "session_" + randomString(8)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[i%len(letters)]
	}
	return string(b)
}
