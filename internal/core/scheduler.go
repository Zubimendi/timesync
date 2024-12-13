package core

import (
	"fmt"
	"time"
)

// Schedule represents a scheduled meeting.
type Schedule struct {
	Time     time.Time
	Timezone string
	Details  string
}

// ScheduleMeeting schedules a meeting and converts it to the default timezone.
func ScheduleMeeting(defaultTZ, inputTZ, details string, t time.Time) (*Schedule, error) {
	loc, err := time.LoadLocation(inputTZ)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone: %s", inputTZ)
	}
	// Convert time to input timezone
	tInInputTZ := t.In(loc)

	// Convert to default timezone
	defaultLoc, err := time.LoadLocation(defaultTZ)
	if err != nil {
		return nil, fmt.Errorf("default timezone is invalid: %s", defaultTZ)
	}
	tInDefaultTZ := tInInputTZ.In(defaultLoc)

	return &Schedule{
		Time:     tInDefaultTZ,
		Timezone: defaultTZ,
		Details:  details,
	}, nil
}
