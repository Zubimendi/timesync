package test

import (
	"testing"
	"time"
	"timesync/internal/core"
	"timesync/config"
)

func TestScheduleMeeting(t *testing.T) {
	cfg := config.LoadConfig()

	tm := time.Date(2023, 10, 10, 14, 0, 0, 0, time.UTC)
	schedule, err := core.ScheduleMeeting(cfg.DefaultTimezone, "Asia/Kolkata", "Team Sync", tm)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if schedule.Time.Location().String() != cfg.DefaultTimezone {
		t.Errorf("Expected timezone %v, got %v", cfg.DefaultTimezone, schedule.Time.Location())
	}
}