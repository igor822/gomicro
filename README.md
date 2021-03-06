Gomicro
=======

This project has the objective of being a skeleton for new micro services projects

# What is this?

Gomicro sets a basic structure of a project using the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).  
With this skeleton you can create new APIs, where it will have [echo](https://echo.labstack.com/) installed already.

## Basic structure

* `appcontext/context.go`: is our `container` of services and object within the project.
* `config/`: defines all necessary configuration that your application needs to run.
* `domain/`: defines all entities of the project
* `gateway/`: defines all connections with the external world, e.g. MongoDB, MySQL
* `usecase/`: defines the use cases and business rules of the application

# Build

```sh
go build
```
