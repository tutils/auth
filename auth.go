package auth

type Auth struct {
	opts Options
}

func NewAuth(opts ...Option) *Auth {
	opt := newOptions(opts...)

	return &Auth{
		opts: *opt,
	}
}
