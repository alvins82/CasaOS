package service

import (
	"testing"

	"github.com/IceWhaleTech/CasaOS/common"
	"github.com/IceWhaleTech/CasaOS/pkg/config"
)

func TestParseReleaseVersion(t *testing.T) {
	tests := []struct {
		name      string
		payload   string
		version   string
		changeLog string
	}{
		{name: "release manifest", payload: `{"version":"v0.4.17-ubuntu26.3","change_log":"Fork updater"}`, version: "v0.4.17-ubuntu26.3", changeLog: "Fork updater"},
		{name: "legacy API", payload: `{"data":{"version":"0.4.15","change_log":"Upstream"}}`, version: "0.4.15", changeLog: "Upstream"},
		{name: "GitHub API", payload: `{"tag_name":"v0.4.17-ubuntu26.3","body":"Release body"}`, version: "v0.4.17-ubuntu26.3", changeLog: "Release body"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := parseReleaseVersion(test.payload)
			if got.Version != test.version || got.ChangeLog != test.changeLog {
				t.Fatalf("parseReleaseVersion() = %#v, want version %q and changelog %q", got, test.version, test.changeLog)
			}
		})
	}
}

func TestResolveUpdateVersionURL(t *testing.T) {
	original := config.ServerInfo.UpdateVersionUrl
	t.Cleanup(func() { config.ServerInfo.UpdateVersionUrl = original })

	config.ServerInfo.UpdateVersionUrl = "https://example.com/version.json"
	if got := resolveUpdateVersionURL(); got != "https://example.com/version.json" {
		t.Fatalf("resolveUpdateVersionURL() = %q", got)
	}

	config.ServerInfo.UpdateVersionUrl = "not a URL"
	if got := resolveUpdateVersionURL(); got != common.FORK_VERSION_URL {
		t.Fatalf("resolveUpdateVersionURL() fallback = %q", got)
	}
}
