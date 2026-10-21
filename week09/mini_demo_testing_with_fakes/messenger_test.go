package minidemotestingwithfakes

import "testing"

type FakeMessenger struct {
	LastMessage string
}

func (f *FakeMessenger) Send(message string) bool {
	f.LastMessage = message
	return true
}

func TestNotify(t *testing.T) {
	fake := &FakeMessenger{}

	ok := Notify(fake, "Riley")

	if !ok {
		t.Fatal("expected Notify to return true")
	}

	if fake.LastMessage != "Hello, Riley" {
		t.Fatalf("got %q", fake.LastMessage)
	}
}
