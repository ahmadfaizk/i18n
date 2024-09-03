module github.com/ahmadfaizk/i18n/examples/go-chi

go 1.18

require (
	github.com/ahmadfaizk/i18n v0.1.0
	github.com/go-chi/chi/v5 v5.1.0
	golang.org/x/text v0.17.0
	gopkg.in/yaml.v3 v3.0.1
)

require github.com/nicksnyder/go-i18n/v2 v2.4.0 // indirect

replace github.com/ahmadfaizk/i18n => ../..
