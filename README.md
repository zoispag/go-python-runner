# Run Python script in Isolation using Golang

This is a Proof of Concect to demonstrate how to run a python script in an isolated environment (virtual environment) via a go wrapper.
It will install dependecies, using one of the following package managers with the following order:

1. Poetry (using `pyproject.toml`)
2. PyFlow (using `pyproject.toml`)
3. pipenv (using `Pipfile`)
4. pip & venv (using `requirements.txt`)

## To install:

### Go requirements

```bash
go mod download
go mod verify
```

### Python package managers

- [Poetry](https://python-poetry.org/docs/#installation)
- [PyFlow](https://github.com/David-OConnor/pyflow#installation)
- [pipenv](https://pipenv-fork.readthedocs.io/en/latest/install.html)

## To run:

```bash
go run .
```

#### with Debug mode:

```bash
GO_DEBUG_MODE=on go run main.go
```

## To build as a binary:

```bash
go build .
```

and execute the produced binary with

```bash
./go-python-runner
```