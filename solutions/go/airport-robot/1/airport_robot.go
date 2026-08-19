package airportrobot

type Greeter interface {
    LanguageName() string
    Greet(name string) string
}

type Italian struct{}

func (Italian) LanguageName() string {
    return "Italian"
}

func (Italian) Greet(name string) string {
    return "Ciao " + name + "!"
}

type Portuguese struct{}

func (Portuguese) LanguageName() string {
    return "Portuguese"
}

func (Portuguese) Greet(name string) string {
    return "Olá " + name + "!"
}

func SayHello(name string, greeter Greeter) string {
    return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}