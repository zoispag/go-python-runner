# AGENTS.md

This file documents the go-python-runner project for AI coding agents.

## Project Overview

`go-python-runner` is a Go library that runs Python scripts in isolated virtual environments. It detects which Python package manager a project uses, creates the virtual environment, and provides the correct Python binary path to execute scripts.

Module path: `github.com/zoispag/go-python-runner`

## Repository Structure

```
go-python-runner/
├── python/               # Core library package
│   ├── virtualenv.go     # Public API: SetupVirtualEnv, CleanUpVirtualEnv
│   ├── run.go            # Public API: GetPythonRunCommand, ExecutePython
│   ├── pep518.go         # Routes pyproject.toml projects to uv or Poetry
│   ├── uv.go             # uv handler: runs `uv sync`
│   ├── poetry.go         # Poetry handler: runs `poetry install`
│   ├── pipenv.go         # pipenv handler: runs `pipenv install`
│   ├── pip.go            # pip+venv handler: creates venv and pip installs
│   ├── utility.go        # fileExists() and dirExists() helpers
│   └── e2e_test.go       # E2E tests (one per package manager)
├── testdata/             # E2E test fixtures (one subdir per package manager)
│   ├── uv/               # uv fixture: pyproject.toml, uv.lock, script.py
│   ├── poetry/           # Poetry fixture: pyproject.toml, poetry.lock, poetry.toml, script.py
│   ├── pipenv/           # pipenv fixture: Pipfile, Pipfile.lock, script.py
│   └── pip/              # pip fixture: requirements.txt, script.py
├── .github/workflows/
│   ├── codeql-analysis.yml
│   └── test.yml          # E2E CI: matrix over uv/poetry/pipenv/pip
├── go.mod
└── go.sum
```

## Public API

All public functions are in the `python` package (`github.com/zoispag/go-python-runner/python`).

### `SetupVirtualEnv(path string)`
Creates a virtual environment at the given path by detecting which package manager to use and running the appropriate install command. See package manager detection below.

### `CleanUpVirtualEnv(path string)`
Deletes the `.venv/` directory if it exists at the given path.

### `ExecutePython(path string, scriptName string) ([]byte, error)`
Runs the given Python script using the virtual environment's Python binary. Returns combined stdout+stderr output.

### `GetPythonRunCommand(path string, scriptName string) *exec.Cmd`
Returns an `*exec.Cmd` for running the Python script, allowing the caller to control execution (Start, Wait, pipes, etc.).

## Package Manager Detection & Precedence

`SetupVirtualEnv` checks for indicator files in this order:

| Priority | Tool | Indicator file | Install command |
|----------|------|---------------|-----------------|
| 1 | **uv** | `uv.lock` (+ `pyproject.toml`) | `uv sync` |
| 2 | **Poetry** | `pyproject.toml` with `[tool.poetry]` section | `poetry install` |
| 3 | **pipenv** | `Pipfile` | `pipenv install` |
| 4 | **pip + venv** | `requirements.txt` | `python3 -m venv .venv` + `pip install -r requirements.txt` |

All supported tools install dependencies into `.venv/` in the project directory. `GetPythonRunCommand` / `ExecutePython` use `.venv/bin/python` if `.venv/` exists, otherwise fall back to system `python`.

## Handler Pattern

Every package manager handler follows the same structure:

```go
func xyzProc(path string) {
    log.Info("Found '<indicator>'. Creating virtual environment using '<Tool>'.")
    cmd := exec.Command("<tool>", "<args>")
    cmd.Dir = path
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Error(fmt.Sprintf("%s", err.Error()))
    }
    log.Debug(string(out))
}
```

Errors are logged via `log.Error`, not returned. This is the established convention — do not change handlers to return errors.

## Build, Test & Run Commands

```bash
go build ./...                              # Build
go vet ./...                                # Vet
go test -v -timeout 120s ./python/          # Run all E2E tests
```

The E2E tests in `python/e2e_test.go` copy each `testdata/<pm>/` fixture into a temp directory, call `SetupVirtualEnv`, run `script.py` via `ExecutePython`, assert an ISO 8601 timestamp is printed, then call `CleanUpVirtualEnv`. Each test skips automatically if the required tool is not on `PATH`.

## Conventions

- Error handling: use `log.Error(fmt.Sprintf(...))` — never return errors from handlers
- New package manager: add `<name>.go` in `python/`, add detection + routing in `pep518.go` or `virtualenv.go`
- Do not modify public function signatures in `virtualenv.go` or `run.go`
- Do not change `ioutil.ReadFile` in `pep518.go` to `os.ReadFile` — leave deprecated API as-is
- Do not fix `fileExists("poetry.toml")` in `poetryProc` — pre-existing CWD-relative behaviour, out of scope

## Commit Rules

- Never add `Co-authored-by` trailers or any co-author lines to commits
