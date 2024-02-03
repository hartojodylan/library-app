# Library App Case Study

For a company :)


- Swagger API documentation generation, environment settings, redis, mysql, mongo
- Use `go mod` to install the management dependencies

# API Usage
note: refer to swagger for body and query parameters
- GET - `/v1/books/{subject}`

    - For get, api will retrieve from OpenLibrary and store the book key `/works/OL66534W` as `OL66534W`
    then book details such as author, title and edition number will be stored in the global variable acting as a database.
    - This API uses paging, so you can specify the limit and page number to retrieve the data.
- POST - `/v1/books`
    
    - For post, a user can book multiple books with one API call. The API will store the book details in the global variable acting as a database.
    If the book already exists, the user won't be able to book said book and will be notified. If the user books multiple books and some of the books are booked, the user will be notified.
    - The user needs to specify the book ids, the user's id and the pick up schedule.

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
core/ Low level logic code directory (such as database connection, cache connection, etc.)
global/ Global variables and constants acting as db
model/  Data and logic code directory
  |- form/  Request form structure data definition, form validation configuration
  |- logic/ Logic processing
  |- mongo/ MongoDB data collection model definition
  |- mysql/ MySQL data form model definition
  |_ rds/   Redis data model definition
resource/   Non-code resources used by some projects (language files, view template files, etc.)
runtime/    Temporary file directory (file cache, log files, etc.)
service/   Service layer directory(high-level logic code, such as business logic, etc.)
static/     Static resource directory (swagger yaml, js, css, etc.)
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
