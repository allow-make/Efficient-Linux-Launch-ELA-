package i18n

import (
	"os"
	"strings"
)

var currentLang = "zh"

func SetLanguage(lang string) {
	if lang == "en" || lang == "zh" {
		currentLang = lang
	}
}

func GetLang() string {
	return currentLang
}

func T(key string) string {
	var dict map[string]string
	if currentLang == "zh" {
		dict = Zh
	} else {
		dict = En
	}
	if val, ok := dict[key]; ok {
		return val
	}
	return key
}

func DetectLang() string {
	lang := os.Getenv("LANG")
	if lang != "" && strings.HasPrefix(lang, "zh") {
		return "zh"
	}
	lang = os.Getenv("LANGUAGE")
	if lang != "" && strings.HasPrefix(lang, "zh") {
		return "zh"
	}
	return "en"
}
