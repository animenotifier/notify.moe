package arn

// SessionID represents a session ID.
type SessionID = ID

// Session stores session-related data.
type Session map[SessionID]interface{}
