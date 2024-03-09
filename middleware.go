package i18n

import (
	"context"
	"net/http"

	"golang.org/x/text/language"
)

type contextKey string

const languageCtxKey contextKey = "i18n-language"

func defaultLanguageExtractFunc(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}

// Middleware is a middleware that sets the language to the context from the request.
//
// It uses the Accept-Language header to get the language. But you can also customize with WithLanguageExtractFunc option.
func Middleware(opts ...MiddlewareOption) func(http.Handler) http.Handler {
	defaultOpts := []MiddlewareOption{
		WithLanguageExtractFunc(defaultLanguageExtractFunc),
	}
	opts = append(defaultOpts, opts...)
	cfg := newMiddlewareConfig(opts...)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lang := cfg.languageExtractFunc(r)
			if lang != "" {
				ctx := context.WithValue(r.Context(), languageCtxKey, lang)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// GetLanguage returns the language tag from the context.
//
// If the language tag is not found, it returns the default language tag.
func GetLanguage(ctx context.Context) language.Tag {
	lang, ok := ctx.Value(languageCtxKey).(string)
	if !ok {
		return language.Und
	}
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil || len(tags) == 0 {
		return language.Und
	}
	return tags[0]
}

// SetLanguageToContext sets the language to the context.
func SetLanguageToContext(ctx context.Context, language string) context.Context {
	return context.WithValue(ctx, languageCtxKey, language)
}
