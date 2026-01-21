package config

import (
	"testing"
	"time"
)

func TestLoad_Defaults(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("ENVIRONMENT", "")
	t.Setenv("PB_URL", "")
	t.Setenv("PB_TIMEOUT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Port != "8080" {
		t.Fatalf("expected port 8080, got %s", cfg.Port)
	}
	if cfg.Environment != "development" {
		t.Fatalf("expected environment development, got %s", cfg.Environment)
	}
	if cfg.PocketBaseTimeout != 5*time.Second {
		t.Fatalf("expected timeout 5s, got %s", cfg.PocketBaseTimeout)
	}
}

func TestLoad_InvalidTimeout(t *testing.T) {
	t.Setenv("PB_TIMEOUT", "invalid")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestLoad_CustomTimeout(t *testing.T) {
	t.Setenv("PB_TIMEOUT", "2s")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.PocketBaseTimeout != 2*time.Second {
		t.Fatalf("expected timeout 2s, got %s", cfg.PocketBaseTimeout)
	}
}
