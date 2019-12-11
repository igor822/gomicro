Gomicro
=======

This project has the objective of being a skeleton for new micro services projects

# What is this?

Gomicro has the purpose to set a basic structure of a project using the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  
With this skeleton you can create new APIs, where it will have [echo](https://echo.labstack.com/) installed already.

## Basic structure

* `appcontext/context.go`: is our `container` of services and object within the project.
* `config/`: here you can define all necessary configuration your application needs to run.
* `domain/`: defines all entities of the project
* `gateway/`: connections with external world, e.g. MongoDB, MySQL
* `usecase/`: for any usecases and business rules of the application we define in this folder

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
