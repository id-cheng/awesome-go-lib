// Command example runs a sample webserver that uses go-i18n/v2/i18n.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var page = template.Must(template.New("").Parse(`
	<!DOCTYPE html>
	<html>
	<body>
	
	<h1>{{.Title}}</h1>
	
	{{range .Paragraphs}}<p>{{.}}</p>{{end}}
	
	</body>
	</html>
`))

func main() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	// No need to load active.en.toml since we are providing default translations.
	// bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.zh.toml")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")
		fmt.Println("lang:", lang)
		fmt.Println("Accept-Language:", accept)
		localizer := i18n.NewLocalizer(bundle, lang, accept)

		name := r.FormValue("name")
		if name == "" {
			name = "Cheng"
		}

		hello := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    "Hello",
				Other: "Hello {{.Name}}",
			},
			TemplateData: map[string]string{
				"Name": name,
			},
		})

		unreadEmailCount, _ := strconv.ParseInt(r.FormValue("unreadEmailCount"), 10, 64)

		myUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "MyUnreadEmails",
				Description: "The number of unread emails I have",
				One:         "I have {{.PluralCount}} unread email.",
				Other:       "I have {{.PluralCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
		})

		personUnreadEmails := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "PersonUnreadEmails",
				Description: "The number of unread emails a person has",
				One:         "{{.Name}} has {{.UnreadEmailCount}} unread email.",
				Other:       "{{.Name}} has {{.UnreadEmailCount}} unread emails.",
			},
			PluralCount: unreadEmailCount,
			TemplateData: map[string]any{
				"Name":             name,
				"UnreadEmailCount": unreadEmailCount,
			},
		})

		err := page.Execute(w, map[string]any{
			"Title": hello,
			"Paragraphs": []string{
				myUnreadEmails,
				personUnreadEmails,
			},
		})
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
