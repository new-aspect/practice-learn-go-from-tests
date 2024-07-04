package main

const (
	french  = "French"
	chinese = "Chinese"

	englishHelloPrefix = "Hello, "
	chineseHelloPrefix = "你好, "
	frenchHelloPrefix  = "Bonjour,"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case chinese:
		return chineseHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
	return
}
