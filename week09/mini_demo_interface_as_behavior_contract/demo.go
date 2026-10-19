package main

import "fmt"

// an interface lists the required methods
type Summarizer interface {
	Summary() string
}

// a struct encapsulates related data
type WeatherReport struct {
	City string
	High int
}

// attach a Summary method to WeatherReport
// WeatherReport satisfies the Summarizer interface because it has a Summary method
func (w WeatherReport) Summary() string {
	return fmt.Sprintf("%s high: %d", w.City, w.High)
}

// PrintSummary only cares that the type has a Summary method
// any type that has Summary (satisfies Summarizer) can be used as the argument
func PrintSummary(s Summarizer) {
	fmt.Println(s.Summary())
}

func main() {
	report := WeatherReport{City: "Austin", High: 91}
	PrintSummary(report)
}
