package i18n

import "net/http"

type middlewareConfig struct {
	languageExtractFunc func(*http.Request) string
}

// MiddlewareOption is the option for the middleware.
type MiddlewareOption func(*middlewareConfig)

func newMiddlewareConfig(opts ...MiddlewareOption) *middlewareConfig {
	config := &middlewareConfig{}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

// WithLanguageExtractFunc sets the language from the request.
//
// Example:
//
//	i18n.Middleware(i18n.WithLanguageExtractFunc(func(r *http.Request) string {
//		lang := r.Header.Get("Accept-Language")
//		if lang == "" {
//			lang = r.URL.Query().Get("lang")
//		}
//		return lang
//	}))
func WithLanguageExtractFunc(languageExtractFunc func(*http.Request) string) MiddlewareOption {
	return func(c *middlewareConfig) {
		c.languageExtractFunc = languageExtractFunc
	}
}
