.DEFAULT_GOAL := help!
.RECIPEPREFIX = >
.PHONY: all usage help clean

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | sed -e 's/: / /'

##Available Commands:

  all:        ## Run all the commands for the entire publishing process
all: route apidoc phar pbimg

  clean:      ## Clean all created artifacts
clean:
	git clean --exclude=.idea/ -fdx

  apidoc:     ## Generate swagger UI document json
apidoc:
	swag init -o static
	rm static/docs.go

  pack:       ## Build and package the application
pack:
	# collect git info to current env config file.
	go build -o ./library-app

  pbprod:     ## Build prod docker image and push to your hub
pbprod:
	go build ./cli/cliapp.go && ./cliapp git
	docker build -f Dockerfile -t your.dockerhub.com/library-app --build-arg app_env=prod .
	docker push your.dockerhub.com/library-app

  pbtest:     ## Build test docker image and push to your hub
pbtest:
	go build ./cli/cliapp.go && ./cliapp git
	docker build -f Dockerfile -t your.dockerhub.com/library-app:test --build-arg app_env=test .
	docker push your.dockerhub.com/library-app:test

  pbaudit:    ## Build audit docker image and push to your hub
pbaudit:
	go build ./cli/cliapp.go && ./cliapp git
	docker build -f Dockerfile -t your.dockerhub.com/library-app:audit --build-arg app_env=audit .
	docker push your.dockerhub.com/library-app:audit

  devimg:     ## Build dev docker image
devimg:
	docker build -f Dockerfile --build-arg app_env=dev -t library-app:dev .

##
##Helper Commands:

  test:   ## Run all the tests
test: fmt lt

  echo:   ## echo test
echo:
	echo hello

  fmt:    ## Run the go fmt
fmt:
	go fmt ./...

  lt:     ## Run the golint tool
lt:
	go lint ./...

  tc:     ## Run the unit tests with code coverage
tc:
	go test -cover ./...
