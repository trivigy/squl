package squl

import (
	"github.com/google/uuid"
)

type session struct {
	uid    uuid.UUID
	Params marker
}

func NewSession(uid uuid.UUID) *session {
	return &session{uid: uid, Params: marker{entries: make(map[string]arg)}}
}
