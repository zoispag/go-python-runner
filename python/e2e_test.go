package python_test

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"testing"

	python "github.com/zoispag/go-python-runner/python"
)

// iso8601Re matches the ISO 8601 datetime string printed by script.py.
var iso8601Re = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}`)

// copyDir recursively copies src into dst (dst must already exist).
func copyDir(t *testing.T, src, dst string) {
	t.Helper()
	err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
	if err != nil {
		t.Fatalf("copyDir %s -> %s: %v", src, dst, err)
	}
}

// fixtureDir returns the absolute path to testdata/<name>.
func fixtureDir(t *testing.T, name string) string {
	t.Helper()
	// This file lives in python/; testdata is one level up.
	abs, err := filepath.Abs(filepath.Join("..", "testdata", name))
	if err != nil {
		t.Fatalf("fixtureDir %s: %v", name, err)
	}
	return abs
}

// requireTool skips the test if the named binary is not on PATH.
func requireTool(t *testing.T, tool string) {
	t.Helper()
	if _, err := exec.LookPath(tool); err != nil {
		t.Skipf("%s not found on PATH — skipping", tool)
	}
}

func runE2E(t *testing.T, pmName string) {
	t.Helper()

	// Copy fixture into a temp dir so .venv doesn't pollute the source tree.
	tmp := t.TempDir()
	copyDir(t, fixtureDir(t, pmName), tmp)

	// SetupVirtualEnv must create .venv.
	python.SetupVirtualEnv(tmp)
	venvPath := filepath.Join(tmp, ".venv")
	if _, err := os.Stat(venvPath); os.IsNotExist(err) {
		t.Fatalf("SetupVirtualEnv did not create .venv for %s", pmName)
	}

	// ExecutePython must run script.py and return an ISO 8601 timestamp.
	out, err := python.ExecutePython(tmp, filepath.Join(tmp, "script.py"))
	if err != nil {
		t.Fatalf("ExecutePython error for %s: %v\noutput: %s", pmName, err, out)
	}
	if !iso8601Re.Match(out) {
		t.Fatalf("unexpected output for %s: %q (want ISO 8601 timestamp)", pmName, out)
	}

	// CleanUpVirtualEnv must remove .venv.
	python.CleanUpVirtualEnv(tmp)
	if _, err := os.Stat(venvPath); !os.IsNotExist(err) {
		t.Fatalf("CleanUpVirtualEnv did not remove .venv for %s", pmName)
	}
}

func TestE2EUv(t *testing.T) {
	requireTool(t, "uv")
	runE2E(t, "uv")
}

func TestE2EPoetry(t *testing.T) {
	requireTool(t, "poetry")
	runE2E(t, "poetry")
}

func TestE2EPipenv(t *testing.T) {
	requireTool(t, "pipenv")
	runE2E(t, "pipenv")
}

func TestE2EPip(t *testing.T) {
	requireTool(t, "python3")
	runE2E(t, "pip")
}
