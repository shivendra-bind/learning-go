package hello

const enHelloPrifix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enHelloPrifix + name
}
