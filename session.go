package squl

import (
	"github.com/google/uuid"
)

// Session represts a single build session. It is an internal object is
// represents the `$` inside of the template. It is exported only for purposes
// of documenting available helper objects.
type Session struct {
	uid    uuid.UUID
	Params Params
}

func newSession(uid uuid.UUID) *Session {
	return &Session{uid: uid, Params: Params{entries: make(map[string]arg)}}
}
