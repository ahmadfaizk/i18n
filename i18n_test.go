package i18n_test

import (
	"context"
	"testing"

	"github.com/ahmadfaizk/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func TestI18n(t *testing.T) {
	t.Run("not initialized", func(t *testing.T) {
		assert.Panics(t, func() {
			i18n.T("message.test")
		})
	})
	err := i18n.Init(language.English,
		i18n.WithUnmarshalFunc("yaml", yaml.Unmarshal),
		i18n.WithTranslationFile("testdata/en.yaml", "testdata/id.yaml"),
	)
	require.NoError(t, err)

	testCases := []struct {
		name            string
		messageID       string
		options         []i18n.LocalizeOption
		expectedMessage string
	}{
		{
			name:            "default",
			messageID:       "message.test",
			expectedMessage: "This is test message",
		},
		{
			name:            "with custom language",
			messageID:       "message.test",
			options:         []i18n.LocalizeOption{i18n.Lang("id")},
			expectedMessage: "Ini adalah pesan tes",
		},
		{
			name:            "with param",
			messageID:       "message.hello",
			options:         []i18n.LocalizeOption{i18n.Param("name", "John")},
			expectedMessage: "Hello, John!",
		},
		{
			name:            "with params",
			messageID:       "message.with_params",
			options:         []i18n.LocalizeOption{i18n.Params(i18n.Map{"param1": "hello", "param2": 123})},
			expectedMessage: "This is message with params: hello and 123",
		},
		{
			name:            "with default message",
			messageID:       "test.with_default_message",
			options:         []i18n.LocalizeOption{i18n.DefaultMessage("This is default message")},
			expectedMessage: "This is default message",
		},
		{
			name:            "not found and use default language",
			messageID:       "message.hello_world",
			options:         []i18n.LocalizeOption{i18n.Lang("id")},
			expectedMessage: "Hello, World!",
		},
		{
			name:            "not found",
			messageID:       "test.not_found",
			expectedMessage: "test.not_found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			message := i18n.T(tc.messageID, tc.options...)
			assert.Equal(t, tc.expectedMessage, message)
		})
	}
}

func TestI18nCtx(t *testing.T) {
	err := i18n.Init(language.English,
		i18n.WithUnmarshalFunc("yaml", yaml.Unmarshal),
		i18n.WithTranslationFile("testdata/en.yaml", "testdata/id.yaml"),
	)
	require.NoError(t, err)

	testCases := []struct {
		name            string
		messageID       string
		options         []i18n.LocalizeOption
		expectedMessage string
		language        string
	}{
		{
			name:            "default",
			messageID:       "message.test",
			expectedMessage: "This is test message",
		},
		{
			name:            "with param",
			messageID:       "message.hello",
			options:         []i18n.LocalizeOption{i18n.Param("name", "John")},
			expectedMessage: "Hello, John!",
		},
		{
			name:            "with params",
			messageID:       "message.with_params",
			options:         []i18n.LocalizeOption{i18n.Params(i18n.Map{"param1": "hello", "param2": 123})},
			expectedMessage: "This is message with params: hello and 123",
		},
		{
			name:            "with default message",
			messageID:       "test.with_default_message",
			options:         []i18n.LocalizeOption{i18n.DefaultMessage("This is default message")},
			expectedMessage: "This is default message",
		},
		{
			name:            "with custom language",
			messageID:       "message.test",
			language:        "id",
			expectedMessage: "Ini adalah pesan tes",
		},
		{
			name:            "not found and use default language",
			messageID:       "message.hello_world",
			language:        "id",
			options:         []i18n.LocalizeOption{i18n.Lang("es")},
			expectedMessage: "Hello, World!",
		},
		{
			name:            "not found",
			messageID:       "test.not_found",
			expectedMessage: "test.not_found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			if tc.language != "" {
				ctx = i18n.SetLanguageToContext(ctx, tc.language)
			}
			message := i18n.TCtx(ctx, tc.messageID, tc.options...)
			assert.Equal(t, tc.expectedMessage, message)
		})
	}
}
