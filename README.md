Gomicro
=======

This project has the objective of being a skeleton for new micro services projects

# What is this?

Gomicro has the purpose to set a basic structure of a project using the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  
With this skeleton you can create new APIs, where it will have [echo](https://echo.labstack.com/) installed already.

## Basic structure

* appcontext/context.go: is our `container` of services and object within the project.
* config/: in here you can define all necessary configuration your application needs to run.
* domain/: `domain` folder you can define all structures (or entites) of the project
* gateway/: it's related with infrastructure and connections with external world
* usecase/: for any usecases and business rules of the application we define in this folder

# Installation

Install go
Install dependencies [dep](https://golang.github.io/dep/docs/installation.html)

## Configure

```sh
dep ensure -update -v
```

## Run

```sh
go build
```
