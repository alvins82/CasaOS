package service

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/IceWhaleTech/CasaOS/pkg/config"
)

func TestIsDebianFamily(t *testing.T) {
	tests := []struct {
		name    string
		release map[string]string
		want    bool
	}{
		{name: "debian", release: map[string]string{"ID": "debian"}, want: true},
		{name: "ubuntu like", release: map[string]string{"ID": "linuxmint", "ID_LIKE": "ubuntu debian"}, want: true},
		{name: "raspberry pi", release: map[string]string{"ID": "raspbian"}, want: true},
		{name: "fedora", release: map[string]string{"ID": "fedora"}, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isDebianFamily(test.release); got != test.want {
				t.Fatalf("isDebianFamily() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestParseAPTUpgradeSimulation(t *testing.T) {
	output := `
NOTE: This is only a simulation!
Inst libc6 [2.36-9+deb12u7] (2.36-9+deb12u8 Debian:12.8/stable [amd64])
Inst linux-image-amd64:amd64 [6.1.128-1] (6.1.129-1 Debian:12.8/stable [amd64])
Conf libc6 (2.36-9+deb12u8 Debian:12.8/stable [amd64])
`

	got := parseAPTUpgradeSimulation(output)
	if len(got) != 2 {
		t.Fatalf("parseAPTUpgradeSimulation() returned %d updates, want 2: %#v", len(got), got)
	}
	if got[0].Name != "libc6" || got[0].CurrentVersion != "2.36-9+deb12u7" || got[0].CandidateVersion != "2.36-9+deb12u8" {
		t.Fatalf("first parsed update = %#v", got[0])
	}
	if got[1].Name != "linux-image-amd64:amd64" {
		t.Fatalf("architecture-qualified package was not preserved: %#v", got[1])
	}
}

func TestParseAPTUpgradeSimulationNoUpdates(t *testing.T) {
	if got := parseAPTUpgradeSimulation("0 upgraded, 0 newly installed, 0 to remove and 0 not upgraded.\n"); len(got) != 0 {
		t.Fatalf("parseAPTUpgradeSimulation() = %#v, want no updates", got)
	}
}

func TestSystemPackageUpdateCommand(t *testing.T) {
	command := systemPackageUpdateCommand("/usr/bin/apt-get", "/var/log/casaos/package update.log")
	for _, expected := range []string{
		"'/usr/bin/apt-get' -y --no-remove -o Dpkg::Use-Pty=0 -o Dpkg::Options::=--force-confold upgrade",
		"CASAOS_PACKAGE_UPDATE_STARTED",
		"CASAOS_PACKAGE_UPDATE_SUCCESS",
		"CASAOS_PACKAGE_UPDATE_FAILED",
		"/var/log/casaos/package update.log",
	} {
		if !strings.Contains(command, expected) {
			t.Fatalf("systemPackageUpdateCommand() does not contain %q: %s", expected, command)
		}
	}
	if strings.Contains(command, "dist-upgrade") || strings.Contains(command, "full-upgrade") {
		t.Fatalf("systemPackageUpdateCommand() uses a non-conservative upgrade: %s", command)
	}
}

func TestSystemPackageUpdateArgsAreDetached(t *testing.T) {
	args := systemPackageUpdateArgs("/usr/bin/apt-get", "/var/log/casaos/package-update.log")
	joined := strings.Join(args, " ")
	if !strings.Contains(joined, "--no-block") || !strings.Contains(joined, "--collect") {
		t.Fatalf("systemPackageUpdateArgs() = %#v, want detached systemd job", args)
	}
}

func TestParseSystemPackageUpdateLog(t *testing.T) {
	log := `CASAOS_PACKAGE_UPDATE_STARTED 2026-08-13T01:02:03Z
CASAOS_PACKAGE_UPDATE_REBOOT_REQUIRED
CASAOS_PACKAGE_UPDATE_SUCCESS 2026-08-13T01:05:06Z
`
	parsed := parseSystemPackageUpdateLog(log)
	if parsed.state != systemPackageUpdateStateSuccess || !parsed.rebootRequired {
		t.Fatalf("parsed success log = %#v", parsed)
	}
	if parsed.startedAt == nil || parsed.completedAt == nil {
		t.Fatalf("timestamps were not parsed: %#v", parsed)
	}
	if parsed.startedAt.Format(time.RFC3339) != "2026-08-13T01:02:03Z" {
		t.Fatalf("started timestamp = %v", parsed.startedAt)
	}

	failed := parseSystemPackageUpdateLog("CASAOS_PACKAGE_UPDATE_FAILED 2026-08-13T01:05:06Z 42\n")
	if failed.state != systemPackageUpdateStateFailed || failed.exitCode == nil || *failed.exitCode != 42 {
		t.Fatalf("parsed failure log = %#v", failed)
	}
}

func TestReadBoundedSystemPackageLog(t *testing.T) {
	path := filepath.Join(t.TempDir(), "package-update.log")
	if err := os.WriteFile(path, []byte(strings.Repeat("0123456789", 10)), 0o644); err != nil {
		t.Fatalf("write test log: %v", err)
	}

	got, err := readBoundedSystemPackageLog(path, 64)
	if err != nil {
		t.Fatalf("readBoundedSystemPackageLog() error = %v", err)
	}
	if len(got) > 64 || !strings.Contains(got, "omitted") {
		t.Fatalf("bounded log = %q, length %d", got, len(got))
	}
}

func TestSystemPackageUpdaterCheck(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	simulation := "Inst zlib1g [1.2.13.dfsg-1] (1.2.13.dfsg-1+deb12u1 Debian:12/stable [amd64])\n"
	var calls []string
	updater.command = func(_ context.Context, name string, args ...string) ([]byte, error) {
		calls = append(calls, name+" "+strings.Join(args, " "))
		if name == "systemctl" {
			return []byte("inactive\n"), nil
		}
		if len(args) > 0 && args[0] == "update" {
			return nil, nil
		}
		return []byte(simulation), nil
	}

	got, err := updater.check()
	if err != nil {
		t.Fatalf("check() error = %v", err)
	}
	if !got.Supported || got.Manager != systemPackageManagerAPT || got.Count != 1 {
		t.Fatalf("check() = %#v", got)
	}
	if len(calls) != 3 || !strings.Contains(calls[1], "apt-get update") || !strings.Contains(calls[2], "upgrade") {
		t.Fatalf("unexpected command sequence: %#v", calls)
	}
}

func TestSystemPackageUpdaterUnsupportedWithoutRoot(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	updater.getEUID = func() int { return 1000 }

	got, err := updater.check()
	if err != nil {
		t.Fatalf("check() error = %v", err)
	}
	if got.Supported || got.Reason == "" || got.Manager != "" {
		t.Fatalf("unsupported check() = %#v", got)
	}
}

func TestSystemPackageUpdaterUnsupportedOSAndMissingAPT(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	updater.readOSRelease = func() (map[string]string, error) {
		return map[string]string{"ID": "fedora"}, nil
	}

	result, err := updater.check()
	if err != nil || result.Supported || result.Reason == "" {
		t.Fatalf("unsupported OS check() = %#v, error = %v", result, err)
	}

	updater.readOSRelease = func() (map[string]string, error) {
		return map[string]string{"ID": "debian"}, nil
	}
	updater.lookPath = func(name string) (string, error) {
		if name == "apt-get" {
			return "", os.ErrNotExist
		}
		return filepath.Join("/usr/bin", name), nil
	}

	result, err = updater.check()
	if err != nil || result.Supported || result.Reason != "apt-get was not found on this host." {
		t.Fatalf("missing apt-get check() = %#v, error = %v", result, err)
	}
}

func TestSystemPackageUpdaterRejectsDuplicateUpdate(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	updater.command = func(_ context.Context, name string, _ ...string) ([]byte, error) {
		if name == "systemctl" {
			return []byte("active\n"), nil
		}
		return nil, nil
	}

	status, err := updater.startUpdate()
	if !errors.Is(err, ErrSystemPackageUpdateRunning) {
		t.Fatalf("startUpdate() error = %v, want ErrSystemPackageUpdateRunning", err)
	}
	if status.State != systemPackageUpdateStateRunning {
		t.Fatalf("duplicate status = %#v", status)
	}
}

func TestSystemPackageUpdaterStartWritesQueuedState(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	var startedPath, startedUnit string
	var startedArgs []string
	updater.start = func(path, unit string, args ...string) ([]byte, error) {
		startedPath = path
		startedUnit = unit
		startedArgs = args
		return nil, nil
	}

	status, err := updater.startUpdate()
	if err != nil {
		t.Fatalf("startUpdate() error = %v", err)
	}
	if status.State != systemPackageUpdateStateRunning || status.StartedAt == nil {
		t.Fatalf("startUpdate() status = %#v", status)
	}
	if startedPath != "/usr/bin/systemd-run" || startedUnit != systemPackageUpdateUnit {
		t.Fatalf("systemd start = %q, %q", startedPath, startedUnit)
	}
	if !strings.Contains(strings.Join(startedArgs, " "), "--no-block") {
		t.Fatalf("systemd args = %#v", startedArgs)
	}
	log, err := os.ReadFile(updater.logPath())
	if err != nil || !strings.Contains(string(log), "CASAOS_PACKAGE_UPDATE_QUEUED") {
		t.Fatalf("queued log = %q, error = %v", log, err)
	}
}

func TestSystemPackageUpdaterStatusParsesCompletionAndReboot(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	log := "CASAOS_PACKAGE_UPDATE_STARTED 2026-08-13T01:02:03Z\nCASAOS_PACKAGE_UPDATE_REBOOT_REQUIRED\nCASAOS_PACKAGE_UPDATE_SUCCESS 2026-08-13T01:05:06Z\n"
	if err := os.WriteFile(updater.logPath(), []byte(log), 0o644); err != nil {
		t.Fatalf("write package update log: %v", err)
	}
	updater.stat = func(path string) (os.FileInfo, error) {
		if path == "/var/run/reboot-required" {
			return &testFileInfo{}, nil
		}
		return os.Stat(path)
	}

	status := updater.status()
	if status.State != systemPackageUpdateStateSuccess || !status.RebootRequired || status.StartedAt == nil || status.CompletedAt == nil {
		t.Fatalf("status() = %#v", status)
	}
}

func TestSystemPackageUpdaterStatusFinalizesBeforeReportingFailure(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	now := updater.now().UTC()
	logPath := updater.logPath()
	log := "CASAOS_PACKAGE_UPDATE_STARTED 2026-08-13T01:02:03Z\n"
	if err := os.WriteFile(logPath, []byte(log), 0o644); err != nil {
		t.Fatalf("write package update log: %v", err)
	}
	updater.stat = func(path string) (os.FileInfo, error) {
		if path == "/var/run/reboot-required" {
			return nil, os.ErrNotExist
		}
		if path == logPath {
			return &testFileInfo{modTime: now.Add(-5 * time.Second)}, nil
		}
		return os.Stat(path)
	}

	status := updater.status()
	if status.State != systemPackageUpdateStateFinalizing || status.Error != "" {
		t.Fatalf("status() during result grace = %#v", status)
	}

	if err := os.WriteFile(logPath, []byte(log+"CASAOS_PACKAGE_UPDATE_SUCCESS 2026-08-13T01:05:06Z\n"), 0o644); err != nil {
		t.Fatalf("write success marker: %v", err)
	}
	status = updater.status()
	if status.State != systemPackageUpdateStateSuccess {
		t.Fatalf("status() after delayed success marker = %#v", status)
	}
}

func TestSystemPackageUpdaterStatusEventuallyReportsMissingResult(t *testing.T) {
	updater := newTestSystemPackageUpdater(t)
	now := updater.now().UTC()
	logPath := updater.logPath()
	if err := os.WriteFile(logPath, []byte("CASAOS_PACKAGE_UPDATE_STARTED 2026-08-13T01:02:03Z\n"), 0o644); err != nil {
		t.Fatalf("write package update log: %v", err)
	}
	updater.stat = func(path string) (os.FileInfo, error) {
		if path == "/var/run/reboot-required" {
			return nil, os.ErrNotExist
		}
		if path == logPath {
			return &testFileInfo{modTime: now.Add(-(systemPackageResultGracePeriod + time.Second))}, nil
		}
		return os.Stat(path)
	}

	status := updater.status()
	if status.State != systemPackageUpdateStateFailed || status.Error == "" {
		t.Fatalf("status() after result grace = %#v", status)
	}
}

type testFileInfo struct {
	modTime time.Time
}

func (*testFileInfo) Name() string      { return "reboot-required" }
func (*testFileInfo) Size() int64       { return 0 }
func (*testFileInfo) Mode() os.FileMode { return 0 }
func (info *testFileInfo) ModTime() time.Time {
	return info.modTime
}
func (*testFileInfo) IsDir() bool      { return false }
func (*testFileInfo) Sys() interface{} { return nil }

func newTestSystemPackageUpdater(t *testing.T) *systemPackageUpdater {
	t.Helper()
	logDir := t.TempDir()
	originalLogPath := config.AppInfo.LogPath
	config.AppInfo.LogPath = logDir
	t.Cleanup(func() { config.AppInfo.LogPath = originalLogPath })

	return &systemPackageUpdater{
		command: func(_ context.Context, name string, _ ...string) ([]byte, error) {
			if name == "systemctl" {
				return []byte("inactive\n"), nil
			}
			return nil, nil
		},
		start:    func(_ string, _ string, _ ...string) ([]byte, error) { return nil, nil },
		lookPath: func(name string) (string, error) { return filepath.Join("/usr/bin", name), nil },
		readOSRelease: func() (map[string]string, error) {
			return map[string]string{"ID": "debian"}, nil
		},
		getEUID:   func() int { return 0 },
		readLog:   readBoundedSystemPackageLog,
		writeFile: os.WriteFile,
		mkdirAll:  os.MkdirAll,
		stat: func(path string) (os.FileInfo, error) {
			if path == "/var/run/reboot-required" {
				return nil, os.ErrNotExist
			}
			return os.Stat(path)
		},
		now: func() time.Time { return time.Date(2026, time.August, 13, 1, 2, 3, 0, time.UTC) },
	}
}
