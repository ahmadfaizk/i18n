# i18n
i18n is a simple internationalization library for Golang.
It provides a way to translate strings into multiple languages.
This library is wrapper for [go-i18n](https://github.com/nicksnyder/go-i18n) with some additional features.

## Installation
```bash
go get -u github.com/ahmadfaizk/i18n
```

## Usage

### Create Translation File
```yaml
# locale/en.yaml
hello: Hello
hello_name: Hello, %s
```

```yaml
# locale/id.yaml
hello: Halo
hello_name: Halo, %s
```

### Basic

```go
package main

import (
	"fmt"
	"github.com/ahmadfaizk/i18n"
	"golang.org/x/text/language"
)

func main() {
	if err := i18n.Init(language.English,
		i18n.WithTranslationFile("locale/en.yaml", "locale/id.yaml"),
	); err != nil {
		panic(err)
	}

	fmt.Println(i18n.T("hello")) // Hello
	fmt.Println(i18n.T("hello", i18n.Lang("id"))) // Halo
	fmt.Println(i18n.T("hello_name", i18n.Param("name", "John"))) // Hello, John
	fmt.Println(i18n.T("hello_name", i18n.Lang("id"), i18n.Param("name", "John"))) // Halo, John
}
```

### Using Context (e.g. with chi router)

```go
package main

import (
	"fmt"
	"github.com/ahmadfaizk/i18n"
	"github.com/go-chi/chi"
	"golang.org/x/text/language"
	"net/http"
)

func main() {
	if err := i18n.Init(language.English,
		i18n.WithTranslationFile("locale/en.yaml", "locale/id.yaml"),
	); err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		message := i18n.TCtx(ctx, "hello")
		_, _ = w.Write([]byte(message))
	})
	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
        name := chi.URLParam(r, "name")
        message := i18n.TCtx(ctx, "hello_name", i18n.Param("name", name))
        _, _ = w.Write([]byte(message))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
        panic(err)
    }
}
```
