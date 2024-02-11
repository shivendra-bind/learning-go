package hello

const (
	enHelloPrifix = "Hello, "
	esHelloPrifix = "Hola, "
	frHelloPrifix = "Bonjour, "

	french  = "French"
	spanish = "Spanish"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frHelloPrifix
	case spanish:
		prefix = esHelloPrifix
	default:
		prefix = enHelloPrifix
	}
	return
}
