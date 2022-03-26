# microservices-skeleton
This repository contains some examples to make microservices in python, golang and typescript, this structure includes dependency injection to manage the different layers as infrastructure, domain, and application.


## Project structure

These microservices are oriented to struct a basic application using layers such as domain, infrastructure, and app, the main goal is to apply dependency injection in the **controllers** because the use cases can not have the responsibility to make instances of objects, by example, infrastructure layer, these dependencies will be injected. 

## Python
In python the skeleton has the following structure.

```
├── app
│   ├── controllers.py   (handle request HTTP and apply dependency injection)
│   ├── di.py            (dependency injection)
│   └── routes.py        (routes and config app)

├── domain
│   └── usecases
│       ├── data_processing_usecase.py

├── infraestructure
│   ├── databases
│   │   ├── mongo.py

│   ├── messaging
│   │   ├── nats.py


├── main.py               (run the app)
```

## Golang
In Golang we have the follow files structure (DI **Dependency Injector**)
```
├── app
│   ├── app.go
│   ├── appRequestHandler.go  (controllers)
│   ├── appRouter.go 
│   ├── providers.go          (providers used for DI)
│   ├── wire_gen.go
│   └── wire.go               (DI to create app)
├── Dockerfile
├── domain
│   ├── gateways.go           (interfaces for infraestructure)
│   ├── helpers
│   │   └── tracing.go        (apm trace)
│   ├── models
│   │   └── dto.go            (structs defined as models)
│   └── usecases
│       └── reportsUsecase.go
├── go.mod                    (manage packages)
├── go.sum   
├── infraestructure
│   ├── cache
│   │   └── cacheImp.go       (redis)
│   └── database
│       └── firestoreImp.go   (firestore)
├── main.go                   (run the app)
```