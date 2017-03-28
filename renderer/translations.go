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

// Translate translates the response in the given language.
func (response *Response) Translate(language string) {
	translate := getTranslateFunction(language)
	response.Data = translateResponse(response.Data, translate)
}

// translateArray translates an array of string given the translate function.
func translateArray(strings []string, translate i18n.TranslateFunc) []string {
	var result []string

	for i := 0; i < len(strings); i++ {
		result = append(result, translate(strings[i]))
	}

	return result
}

// translateResponse translates the whole response recursively given the
// translate function.
func translateResponse(response map[string]interface{}, translateFunction i18n.TranslateFunc) map[string]interface{} {
	for key, value := range response {
		switch value.(type) {
		case string:
			response[key] = translateFunction(value.(string))
		case []string:
			response[key] = translateArray(value.([]string), translateFunction)
		case map[string]interface{}:
			response[key] = translateResponse(value.(map[string]interface{}), translateFunction)
		}
	}

	return response
}
