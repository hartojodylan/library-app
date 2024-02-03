# Library App Case Study

For a company :)


- Swagger API documentation generation, environment settings, redis, mysql, mongo
- Use `go mod` to install the management dependencies

## Project Structure

```text
api/ API interface application handlers
  |- controller/
  |- middleware/
  |_ routes.go
app/ Common directory (public methods, application initialization, public components, etc.)
cmd/ CLI command line application commands
  |_ cliapp/ command line application entry file (main)
config/ Application configuration directory (basic configuration plus various environment configurations)
model/  Data and logic code directory
  |- form/  Request form structure data definition, form validation configuration
  |- logic/ Logic processing
  |- mongo/ MongoDB data collection model definition
  |- mysql/ MySQL data form model definition
  |_ rds/   Redis data model definition
resource/   Non-code resources used by some projects (language files, view template files, etc.)
runtime/    Temporary file directory (file cache, log files, etc.)
static/     Static resource directory (js, css, etc.)
main.go     Web application entry file
Dockerfile  Dockerfile
Makefile    Has written some common shortcut commands to help package, build docker, generate documentation, run tests, etc.
...
```

> render by `tree -d 2 ./`

## Start
- Run `go mod tidy` to install dependent libraries
- Run the project: `go run main.go`
- To install swagger use: `go install github.com/swaggo/swag/cmd/swag@latest`

### Init project

```shell
go run ./cmd/appinit
```

### Swagger Docs Generation

installation:

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

> Please check the documentation and examples of `swaggo/swag`

Generated to the specified directory:

```bash
swag init -o static
# This file will be generated at the same time. It can be deleted if it is not needed.
rm static/docs.go
```

Notice:

> `swaggo/swag` is the parsing field description information from the comment of the field

```go
type SomeModel struct {
	// the name description
	Name   string `json:"name" example:"tom"`
}	
```

## Help

- Run the test

```bash
Go test
// output coverage
Go test -cover
```

- Formatting project

```bash
gofmt -s -l ./
go fmt ./...
```

- Run GoLint check

> Note: You need to install `GoLint` first.

```bash
golint ./...
```
