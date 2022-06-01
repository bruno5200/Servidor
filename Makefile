# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=server
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WIN=$(BINARY_NAME)_win.exe
BINARY_OSX=$(BINARY_NAME)_osx
EXPOSE_PORT=127.0.0.1:3000:3000
REPO_NAME=tobias0406/$(BINARY_NAME)
TAG=latest

#ejecuta la compilación con el modo verbose
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# ejecuta test
test:
	$(GOTEST) -v ./...

#descarga las dependencias
mod:
	$(GOMOD) tidy

run-bin:
#$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

#limpia los binarios compilados
clean:
	$(GOCLEAN)
	rm -f $(BINARY_UNIX)

# ejecuta el main.go
run:
	$(GORUN) main.go
	
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

#hace el build de la imagen Docker
docker-build:
	docker build -t $(REPO_NAME):latest .

#hace el run de la imagen Docker
docker-run:
	docker run -p $(EXPOSE_PORT) $(REPO_NAME):latest

#hace el push de la imagen Docker
docker-push:
	docker push $(REPO_NAME):latest