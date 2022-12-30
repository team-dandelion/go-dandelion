package http

import routing "github.com/gly-hub/fasthttp-routing"

type (
	Option  func(*options)
	options struct {
		resourceExtract func(*routing.Context) string
		blockFallback   func(*routing.Context) error
	}
)

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	for _, opt := range opts {
		opt(optCopy)
	}

	return optCopy
}

// WithResourceExtractor sets the resource extractor of the web requests.
func WithResourceExtractor(fn func(*routing.Context) string) Option {
	return func(opts *options) {
		opts.resourceExtract = fn
	}
}

// WithBlockFallback sets the fallback handler when requests are blocked.
func WithBlockFallback(fn func(ctx *routing.Context) error) Option {
	return func(opts *options) {
		opts.blockFallback = fn
	}
}
