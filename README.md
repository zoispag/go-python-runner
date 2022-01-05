# Go-Python-Runner <img src="https://user-images.githubusercontent.com/21138205/84591786-7d81fe80-ae41-11ea-9839-5ec82d809178.png" width="40" height="40" alt=":go-python-runner:" class="emoji" title=":go-python-runner:"/> [![GoDoc](https://godoc.org/github.com/zoispag/go-python-runner?status.svg)](https://godoc.org/github.com//zoispag/go-python-runner)

**Go-Python-Runner** is a library that runs Python code in isolation.

It runs a python script in an isolated environment (virtual environment) via a go wrapper.
It will install dependecies, using one of the following package managers with the following order:

1. Poetry (using `pyproject.toml`)
2. PyFlow (using `pyproject.toml`)
3. pipenv (using `Pipfile`)
4. pip & venv (using `requirements.txt`)

### Usage

#### Create virtual environment

```go
python.SetupVirtualEnv("/path/to/python/script/")
```

This will create a vitual enviroment given the existence of proper files (`pyproject.toml`, `Pipfile` or `requirements.txt`).

#### Cleanup virtual environment

```go
python.CleanUpVirtualEnv("/path/to/python/script/")
```

This will delete the `.venv` directory if exists. In case of pyflow, it will also delete the `__pypackages__` directory.

#### Run python script inside virtual environment

```go
out, err := python.ExecutePython("/path/to/python/script/", "/path/to/python/script/script.py")
```

This will run a python script called `script.py` inside the virtual environment, using the proper command, analyzing files existence.

Alternatively, it is possible to get an instance for `*exec.Cmd` which can be handled independently.

```go
	cmd := python.GetPythonRunCommand("/path/to/python/script/", "/path/to/python/script/script.py")
	out, err := cmd.CombinedOutput()
```
or
```go
	cmd := python.GetPythonRunCommand("/path/to/python/script/", "/path/to/python/script/script.py")
	err := cmd.Run()
```
or
```go
	cmd := python.GetPythonRunCommand("/path/to/python/script/", "/path/to/python/script/script.py")
	err := cmd.Start()
	err = cmd.Wait()
```

#### Complete example

```go
package main

import (
	"fmt"
	"os"

	python "github.com/zoispag/go-python-runner/python"
)

func main() {
	python.SetupVirtualEnv("/path/to/python/script/")

	out, err := python.ExecutePython("/path/to/python/script/", "/path/to/python/script/script.py")
	if err != nil {
		log.Error(fmt.Sprintf("Encountered error: %s", err.Error()))
	}
	log.Info(string(out))
}
```

> **Note**: go-python-runner will **not** install the python package managers. See below for installation instructions.

### Python package managers

- [Poetry](https://python-poetry.org/docs/#installation)
- [PyFlow](https://github.com/David-OConnor/pyflow#installation)
- [pipenv](https://pipenv-fork.readthedocs.io/en/latest/install.html)

### Contribute

#### To run:

```bash
go run .
```

##### with Debug mode:

```bash
GO_DEBUG_MODE=on go run main.go
```

#### To build as a binary:

```bash
go build .
```

and execute the produced binary with

```bash
./go-python-runner
```