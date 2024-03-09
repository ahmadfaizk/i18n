package i18n

import (
	"context"
	"embed"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type translationFSFile struct {
	fs    embed.FS
	paths []string
}

type config struct {
	unmarshalFuncMap          map[string]i18n.UnmarshalFunc
	translationFiles          []string
	translationFSFiles        []translationFSFile
	extractLanguageFunc       func(ctx context.Context) string
	missingTranslationHandler func(id string, err error) string
}

// Option is the option for the i18n package.
type Option func(*config)

func newI18nConfig(opts ...Option) *config {
	c := &config{
		unmarshalFuncMap: make(map[string]i18n.UnmarshalFunc),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithUnmarshalFunc sets the unmarshal function for the bundle.
//
// It is used to unmarshal the message file. You can use yaml.Unmarshal, json.Unmarshal, or any other unmarshal function.
func WithUnmarshalFunc(format string, unmarshalFunc i18n.UnmarshalFunc) Option {
	return func(c *config) {
		c.unmarshalFuncMap[format] = unmarshalFunc
	}
}

// WithTranslationFile sets the message file paths for the bundle.
func WithTranslationFile(paths ...string) Option {
	return func(c *config) {
		c.translationFiles = append(c.translationFiles, paths...)
	}
}

// WithTranslationFSFile sets the message file paths for the bundle.
//
// It is similar to WithTranslationFile, but it uses embed.FS as file system.
func WithTranslationFSFile(fs embed.FS, paths ...string) Option {
	return func(c *config) {
		c.translationFSFiles = append(c.translationFSFiles, translationFSFile{fs: fs, paths: paths})
	}
}

// WithMissingTranslationHandler sets the missing translation handler for the bundle.
//
// It is used to handle the missing translation. The default handler returns the message ID.
func WithMissingTranslationHandler(missingTranslationHandler func(id string, err error) string) Option {
	return func(c *config) {
		c.missingTranslationHandler = missingTranslationHandler
	}
}

// WithExtractLanguageFunc sets the language extract function for the middleware.
//
// It is used to extract the language from the context.
func WithExtractLanguageFunc(extractLanguageFunc func(ctx context.Context) string) Option {
	return func(c *config) {
		c.extractLanguageFunc = extractLanguageFunc
	}
}
