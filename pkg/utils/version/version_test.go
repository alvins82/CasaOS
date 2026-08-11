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
		{name: "next bundle", latest: "v0.4.17-ubuntu26.3", current: "v0.4.17-ubuntu26.2", want: true},
		{name: "double digit bundle", latest: "v0.4.17-ubuntu26.10", current: "v0.4.17-ubuntu26.9", want: true},
		{name: "next CasaOS version", latest: "v0.4.18-ubuntu26.1", current: "v0.4.17-ubuntu26.9", want: true},
		{name: "same bundle", latest: "v0.4.17-ubuntu26.3", current: "v0.4.17-ubuntu26.3", want: false},
		{name: "older bundle", latest: "v0.4.17-ubuntu26.2", current: "v0.4.17-ubuntu26.3", want: false},
		{name: "missing latest", latest: "", current: "v0.4.17-ubuntu26.3", want: false},
		{name: "invalid latest", latest: "latest", current: "v0.4.17-ubuntu26.3", want: false},
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

	if err := os.WriteFile(versionFile, []byte("v0.4.17-ubuntu26.4\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if got := currentVersionFromFile(versionFile); got != "v0.4.17-ubuntu26.4" {
		t.Fatalf("installed version returned %q", got)
	}
}
