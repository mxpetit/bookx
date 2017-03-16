package renderer

import (
	"github.com/nicksnyder/go-i18n/i18n"
)

// getTranslateFunction returns the translation function for the given
// language.
func getTranslateFunction(language string) i18n.TranslateFunc {
	T, _ := i18n.Tfunc(language)

	return T
}

// Translate translates the response's message in the given language.
func (response *Response) Translate(language string) {
	translate := getTranslateFunction(language)

	for key, value := range response.Data {
		switch value.(type) {
		case string:
			response.Data[key] = translate(value.(string))
		}
	}
}
