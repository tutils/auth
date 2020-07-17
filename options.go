package auth

import (
	"auth/provider"
	"auth/sessionstorer"
	"auth/userstorer"
)

type Options struct {
	provider      provider.Provider
	userStorer    userstorer.UserStorer
	sessionStorer sessionstorer.SessionStorer
}

type Option func(opts *Options)

func newOptions(opts ...Option) *Options {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}

	// set default options here

	return opt
}

func WithProvider(provider provider.Provider) Option {
	return func(opts *Options) {
		opts.provider = provider
	}
}

func WithUserStorer(userStorer userstorer.UserStorer) Option {
	return func(opts *Options) {
		opts.userStorer = userStorer
	}
}

func WithSessionStorer(sessionStorer sessionstorer.SessionStorer) Option {
	return func(opts *Options) {
		opts.sessionStorer = sessionStorer
	}
}
