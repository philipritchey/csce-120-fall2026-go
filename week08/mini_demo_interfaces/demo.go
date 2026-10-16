package main

import "fmt"

// Interfaces are useful when you want to write code that works with different kinds of values,
//   as long as those values can do the same required behavior.

// Interfaces are like types because you can use them in places where you would normally use a type

// A struct type is about what data something stores
// An interface type is about what behavior something provides

// An interface describes what methods a type must have

type Alertable interface {
	Alert() string
}

// Anything with an Alert() string method satisfies Alertable
// Anything which satisfies Alertable can be used as an Alertable value

// sendAlert can accept any value whose type has an Alert() method.

func sendAlert(item Alertable) {
	fmt.Println(item.Alert())
}

// Two different struct types with different fields
type AlarmClock struct {
	Time string
}

type PhoneReminder struct {
	Message string
}

// Both types have an Alert method, so both satisfy Alertable.

func (a AlarmClock) Alert() string {
	return "Alarm ringing at " + a.Time
}

func (p PhoneReminder) Alert() string {
	return "Reminder: " + p.Message
}

// There is no need to explicitly state that we intend AlarmClock or PhoneReminder
// to implement Alertable.
// A type satisfies the interface if it has the required method.
// The compiler will check that values passed to sendAlert implement Alertable automatically.

func main() {
	clock := AlarmClock{Time: "7:30 AM"}
	reminder := PhoneReminder{Message: "Submit lab"}

	// clock and reminder are different concrete types,
	// but both satisfy the same interface

	sendAlert(clock)
	sendAlert(reminder)
}

// An interface is a type that describes what methods a value must have,
// rather than what fields it stores.
