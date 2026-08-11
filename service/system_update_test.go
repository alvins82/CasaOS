package service

import (
	"testing"

	"github.com/IceWhaleTech/CasaOS/common"
	"github.com/IceWhaleTech/CasaOS/pkg/config"
)

func TestResolveUpdateInstallerURL(t *testing.T) {
	original := config.ServerInfo.UpdateUrl
	t.Cleanup(func() { config.ServerInfo.UpdateUrl = original })

	config.ServerInfo.UpdateUrl = "https://example.com/install.sh"
	if got := resolveUpdateInstallerURL(); got != "https://example.com/install.sh" {
		t.Fatalf("resolveUpdateInstallerURL() = %q", got)
	}

	config.ServerInfo.UpdateUrl = "http://example.com/install.sh"
	if got := resolveUpdateInstallerURL(); got != common.FORK_UPDATE_URL {
		t.Fatalf("resolveUpdateInstallerURL() fallback = %q", got)
	}
}

func TestShellQuote(t *testing.T) {
	got := shellQuote("https://example.com/a'b")
	want := `'https://example.com/a'"'"'b'`
	if got != want {
		t.Fatalf("shellQuote() = %q, want %q", got, want)
	}
}
