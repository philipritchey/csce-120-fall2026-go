package minidemotestingwithfakes

import "testing"

// A fake type ("mock") that can stand-in for a Messenger and "spy" on it
type FakeMessenger struct {
	LastMessage string
}

// FakeMessenger has Send(), so satisfies the Messenger interface
// Instead of actually sending the message, we just log it and return success
//
//	this keeps the test isolated (no external dependencies)
//	and allows us to check that the correct message was sent, e.g. by Notify()
func (f *FakeMessenger) Send(message string) bool {
	f.LastMessage = message
	return true
}

func TestNotify(t *testing.T) {
	// make a fake messenger to spy on calls to Send(), e.g. by Notify())
	fake := &FakeMessenger{}

	// use the fake messenger to send a message to Riley
	ok := Notify(fake, "Riley")

	if !ok {
		t.Fatal("expected Notify to return true")
	}

	// make sure Notify() did the right thing
	if fake.LastMessage != "Hello, Riley" {
		t.Fatalf("got %q", fake.LastMessage)
	}

	// This should remind you of using bytes.Buffer or strings.Builder
	// to capture output, e.g. redirecting from Stdout to an io.Writer interface
}
