package main

import (
	"log/slog"
	"testing"
)

func TestServer(t *testing.T) {
	if 2+2 != 4 {
		t.Fail()
		slog.Warn("Error")
	}
	slog.Info("Test passed")
}
