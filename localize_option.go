package i18n

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Map is a map of string to interface. It is used to pass template data to the message.
type Map map[string]interface{}

type localizeConfig struct {
	params         Map
	defaultMessage string
	language       string
}

func newLocalizeConfig(opts ...LocalizeOption) *localizeConfig {
	c := &localizeConfig{
		params: make(Map),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c localizeConfig) toI18nLocalizeConfig(id string) *i18n.LocalizeConfig {
	localizeConfig := &i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: c.params,
	}
	if c.defaultMessage != "" {
		localizeConfig.DefaultMessage = &i18n.Message{
			ID:    id,
			Other: c.defaultMessage,
		}
	}
	return localizeConfig
}

// LocalizeOption is a function that configures the localizeConfig.
type LocalizeOption func(*localizeConfig)

// Params sets template data for the message.
//
// Example:
//
//	i18n.T("hello", i18n.Params(i18n.Map{"name": "John"}))
func Params(params map[string]interface{}) LocalizeOption {
	return func(c *localizeConfig) {
		for k, v := range params {
			c.params[k] = v
		}
	}
}

// Param sets template data for the message.
//
// Example:
//
//	i18n.T("hello", i18n.Param("name", "John"))
func Param(key string, value interface{}) LocalizeOption {
	return func(c *localizeConfig) {
		c.params[key] = value
	}
}

// Lang sets the language for the message.
//
// Example:
//
//	i18n.T("hello", i18n.Lang("id"))
func Lang(language string) LocalizeOption {
	return func(c *localizeConfig) {
		c.language = language
	}
}

// DefaultMessage sets the default message for the message.
//
// It is used when the message is not found.
//
// Example:
//
//	i18n.T("hello", i18n.DefaultMessage("Hello, {{.name}}!"), i18n.Param("name", "John")))
func DefaultMessage(defaultMessage string) LocalizeOption {
	return func(c *localizeConfig) {
		c.defaultMessage = defaultMessage
	}
}
