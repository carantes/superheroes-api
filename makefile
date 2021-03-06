# Go parameters
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DOCKERCMD=docker-compose
BINARY_NAME=bin/superheroesapi

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v ./... --cover

run:
	$(GORUN) main.go

run-docker:
	${DOCKERCMD} up --build

all: test run
