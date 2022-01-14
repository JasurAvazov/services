package contracts

import "context"

type Session interface {
	Close(err error)
}

type ISessionProvider interface {
	StartSession(c context.Context) (s Session, ctx context.Context, err error)
}
