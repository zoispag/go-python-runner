# Run Python script in Isolation using Golang

This is a Proof of Concect to demonstrate how to run a python script in an isolated environment (virtual environment) via a go wrapper.
It will install dependecies, using one of the following package managers with the following order:

1. Poetry (using `pyproject.toml`)
2. pipenv (using `Pipfile`)
3. pip & venv (using `requirements.txt`)

## To install:

### Go requirements

```bash
go mod download
go mod verify
```

### Python package managers

- [Poetry](https://python-poetry.org/docs/#installation)
- [pipenv](https://pipenv-fork.readthedocs.io/en/latest/install.html)

## To run:

```bash
go run main.go
```