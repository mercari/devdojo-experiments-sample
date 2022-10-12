package main

// This sample is for learning purposes during DevDojo.

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	tutorial := &Tutorial{
		eventLogger: &EventLogger{},
	}
	experiments := GetExperiments("Experiment_1", "Experiment_2", "Experiment_3")

	variant := experiments["Experiment_2"] // *** What is wrong here?
	tutorial.Show(variant)
}

// Tutorial decides which version of the tutorial to show if any.
type Tutorial struct {
	eventLogger *EventLogger
}

func (t *Tutorial) Show(variant uint) {
	t.eventLogger.LogEvent("Experiment_2", fmt.Sprintf("variant:%d", variant))
	if variant == 0 || variant == 1 {
		// don't show anything for users not assigned to the experiment (0) and
		// the control group (1)
		return
	}
	if variant == 2 {
		t.Slideshow()
		return
	}
	if variant == 3 {
		t.Animation()
		return
	}
	if variant == 4 {
		t.DemoVideo()
		return
	}
	// *** On-hands session: How can we improve maintainability of this code?
}

func (t *Tutorial) Slideshow() {
	// omitted for brevity
}

func (t *Tutorial) Animation() {
	// omitted for brevity
}

func (t *Tutorial) DemoVideo() {
	// omitted for brevity
}

// Results maps experiment names to their assigned variants.
type Results map[string]uint

func GetExperiments(names ...string) Results {
	// *** What is wrong with this function implementation???
	var response *http.Response
	var err error
	const maxRetries = 10
	for c := 0; c < maxRetries; c++ {
		response, err = http.Get(generateURL(names))
		if err != nil {
			log.Println("failed to fetch experiments")
		}
		if err == nil {
			break
		}
	}
	if err != nil {
		return Results{}
	}
	return parseExperiments(response)
}

func parseExperiments(response *http.Response) Results {
	// omitted for brevity
	return Results{}
}

func generateURL(names []string) string {
	param := url.QueryEscape(strings.Join(names, ","))
	return fmt.Sprintf("https://localhost:55555/experiments?names=%s", param)
}

type EventLogger struct{}

func (el *EventLogger) LogEvent(details ...string) {
	// Omitted for brevity. Assume all relevant details are properly logged.
}
