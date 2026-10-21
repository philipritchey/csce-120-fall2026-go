package minidemotestingwithfakes

type Messenger interface {
	Send(message string) bool
}

func Notify(m Messenger, name string) bool {
	return m.Send("Hello, " + name)
}
