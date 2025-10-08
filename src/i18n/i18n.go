package i18n

import (
	"embed"
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Embed translation JSON files into the binary
//go:embed lang/active.en.json
//go:embed lang/active.fr.json
//go:embed lang/active.de.json
//go:embed lang/active.es.json
//go:embed lang/active.it.json
var embeddedLangFiles embed.FS

var localizer *i18n.Localizer
var currentLang string

// InitI18n loads all language files and sets the active locale
func InitI18n(lang string) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	files := []string{
		"lang/active.en.json",
		"lang/active.fr.json",
		"lang/active.de.json",
		"lang/active.es.json",
		"lang/active.it.json",
	}
	for _, f := range files {
		data, _ := embeddedLangFiles.ReadFile(f)
		bundle.MustParseMessageFileBytes(data, f)
	}
	if lang != "fr" && lang != "en" && lang != "de" && lang != "es" && lang != "it" {
		lang = "en"
	}
	localizer = i18n.NewLocalizer(bundle, lang)
	currentLang = lang
}

// T returns the localized string for a given message ID
func T(id string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: id})
}
