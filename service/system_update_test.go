package service

import (
	"reflect"
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

func TestDetachedUpdateCommand(t *testing.T) {
	got := detachedUpdateCommand("https://example.com/a'b", "/var/log/casaos/upgrade.log")
	want := `set -o pipefail; exec >> '/var/log/casaos/upgrade.log' 2>&1; curl -fsSL 'https://example.com/a'"'"'b' | /bin/bash`
	if got != want {
		t.Fatalf("detachedUpdateCommand() = %q, want %q", got, want)
	}
}

func TestDetachedUpdateArgs(t *testing.T) {
	got := detachedUpdateArgs("https://example.com/install.sh", "/var/log/casaos/upgrade.log", "v0.4.19")
	want := []string{
		"--quiet",
		"--collect",
		"--unit=casaos-update",
		"--property=Type=exec",
		"--setenv=CASAOS_INSTALLER_DETACHED=1",
		"--description=CasaOS update v0.4.19",
		"/bin/bash",
		"-c",
		"set -o pipefail; exec >> '/var/log/casaos/upgrade.log' 2>&1; curl -fsSL 'https://example.com/install.sh' | /bin/bash",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("detachedUpdateArgs() = %#v, want %#v", got, want)
	}
}
