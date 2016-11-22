package main

import (
	"strconv"

	"github.com/coreos/go-systemd/sdjournal"
)

func AddLogFilters(journal *sdjournal.Journal, config *Config) {
	// Add Priority Filters
	if config.LogPriority < DEBUG {
		for p := 0; p <= int(config.LogPriority); p++ {
			journal.AddMatch("PRIORITY=" + strconv.Itoa(p))
		}
		journal.AddDisjunction()
	}
}
