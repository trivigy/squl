package squl

import (
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type context struct {
	sessions sync.Map
}

func (r *context) Open() (*session, error) {
	sess := NewSession(uuid.New())
	if _, loaded := r.sessions.LoadOrStore(sess.uid.String(), sess); loaded {
		return nil, errors.Errorf("sess uuid conflict %q", sess.uid.String())
	}
	return sess, nil
}

func (r *context) Close(sess *session) string {
	r.sessions.Delete(sess.uid.String())
	return ""
}
