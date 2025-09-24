# Project Structure

The project structure is based on the unofficial standard defined by
[golang-standards](https://github.com/golang-standards) on GitHub.
The repository can be found
[here](https://github.com/golang-standards/project-layout).

However, here's an overview over the structure of this repository.

## Applications (`cmd/`)

The executables of this projects can be found in the `cmd/`
directory. They handle interface specific logic and make use of
internal or external libraries or packages.

## Internal Packages (`internal/`)

Packages that are used in the project internally and that are not
meant for external use, are located in the `internal/` directory.
They are modules of project-relevant functions.

## Public Packages (`pkg/`)

Packages meant to be used by others are located in the `pkg/`
directory.

## API Documentation (`api/`)

The files for documentation of the API are located in the `api/`
directory. It also contains everything related to the generation of
additional documentation or client libraries.

## Unit Tests (`*_test.go`)

Files that end with `_test.go` that are not in the `test/` directory,
specify unit tests for internal components. They can be run using
`make test`.

## Integration Tests (`test/`)

The files within the `test/` directory are used to conduct
integration tests. 

The tests that evaluate the interfaces to other systems are located
in the `test/` directory. They are run using `make test`.

