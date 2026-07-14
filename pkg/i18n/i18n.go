package i18n

var currentLang = "zh"

func SetLanguage(lang string) {
currentLang = lang
}

func T(key string) string {
if currentLang == "zh" {
if val, ok := Zh[key]; ok {
return val
}
}
return key
}
