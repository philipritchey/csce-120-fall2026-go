package minidemotestingwithfakes

// Messengers can Send(string)
type Messenger interface {
	Send(message string) bool
}

// Notify sends a message using Messenger m: "Hello, <name>"
func Notify(m Messenger, name string) bool {
	return m.Send("Hello, " + name)
}
