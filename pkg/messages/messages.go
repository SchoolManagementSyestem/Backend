package messages

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var translations = make(map[string]map[string]string)

func init() {
	loadTranslations("en", "config/locales/en.json")
	loadTranslations("bn", "config/locales/bn.json")
}

// loadTranslations loads translations from a JSON file into the translations map
func loadTranslations(lang, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Could not load translations for language %s: %v", lang, err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	langTranslations := make(map[string]string)
	if err := decoder.Decode(&langTranslations); err != nil {
		log.Printf("Error decoding translations for language %s: %v", lang, err)
		return
	}

	translations[lang] = langTranslations
}

// GetMessage retrieves the message for a given language and key
func GetMessage(lang, key string) (string, bool) {
	if messages, ok := translations[lang]; ok {
		if msg, exists := messages[key]; exists {
			return msg, true
		}
	}
	return "", false
}

// GetMessageOrDefault retrieves the message for a given language and key, or returns a default message in English
func GetMessageOrDefault(lang, key string) string {
	if msg, ok := GetMessage(lang, key); ok {
		return msg
	}
	return translations["en"][key]
}

// GetLangFromRequest retrieves the language from the request headers or defaults to English
func GetLangFromRequest(r *http.Request) string {
	lang := r.Header.Get("Accept-Language")
	if lang == "" || translations[lang] == nil {
		lang = "en" // Default to English if no valid language is provided
	}
	return lang
}
