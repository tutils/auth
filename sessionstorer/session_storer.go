package sessionstorer

type SessionStorer interface {
	Get(uid int64) (*Session, error)
	Put(uid int64, sess *Session) error
}
