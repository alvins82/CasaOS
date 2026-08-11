package version

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/IceWhaleTech/CasaOS/common"
)

func TestIsVersionNewer(t *testing.T) {
	tests := []struct {
		name    string
		latest  string
		current string
		want    bool
	}{
		{name: "platform-neutral migration", latest: "v0.4.18", current: "v0.4.17-ubuntu26.3", want: true},
		{name: "next release", latest: "v0.4.19", current: "v0.4.18", want: true},
		{name: "next CasaOS minor", latest: "v0.5.0", current: "v0.4.18", want: true},
		{name: "same release", latest: "v0.4.18", current: "v0.4.18", want: false},
		{name: "older compatibility release", latest: "v0.4.17-ubuntu26.3", current: "v0.4.18", want: false},
		{name: "missing latest", latest: "", current: "v0.4.18", want: false},
		{name: "invalid latest", latest: "latest", current: "v0.4.18", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := IsVersionNewer(test.latest, test.current); got != test.want {
				t.Fatalf("IsVersionNewer(%q, %q) = %v, want %v", test.latest, test.current, got, test.want)
			}
		})
	}
}

func TestCurrentVersionFromFile(t *testing.T) {
	versionFile := filepath.Join(t.TempDir(), "fork-release")
	if got := currentVersionFromFile(versionFile); got != common.FORK_RELEASE_VERSION {
		t.Fatalf("missing version file returned %q", got)
	}

	if err := os.WriteFile(versionFile, []byte("v0.4.19\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if got := currentVersionFromFile(versionFile); got != "v0.4.19" {
		t.Fatalf("installed version returned %q", got)
	}
}
