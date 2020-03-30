package main

import (
	"os"

	"github.com/getsentry/sentry-go"
)

func sentryInit() bool {
	// Read some configuration values from environment variables
	// (they were loaded from the ".env" file in "main.go")
	sentryDSN := os.Getenv("SENTRY_DSN")
	if len(sentryDSN) == 0 {
		logger.Info("The \"sentryDSN\" environment variable is blank; " +
			"aborting Sentry initialization.")
		return false
	}

	// Initialize Sentry
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	}); err != nil {
		logger.Fatal("Failed to initialize Sentry:", err)
	}

	return true
}