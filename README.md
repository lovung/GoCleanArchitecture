# Golang Clean Architecture template
## Clean Architecture blog post
[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
![Layers](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

## Skeleton
```
.
├── app                     The place of all business logic of the application
│   ├── domain                  Domain layer: enterprise business logic rules
│   │   ├── entity                  All of entities in the application
│   │   └── repository              The interfaces to impact the entities
│   ├── external                External layer: frameworks & drivers
│   │   └── adaptor                 Client to call other services/endpoints
│   ├── interface               The boundary of our data and business logic
│   │   ├── grpc                    Interface to serve the gRPC API
│   │   ├── persistence             Client to call all of persistences
│   │   │   ├── fs                      File system client
│   │   │   ├── kv                      Key-Value client
│   │   │   ├── nosql                   NoSQL client
│   │   │   └── rdbms                   SQL client
│   │   └── restful                 Interface to serve RESTful API
│   │       ├── handler                 (Or controller) Endpoint to serve the RESTful call
│   │       ├── middleware              Middleware for RESTful APIs
│   │       └── presenter               Represent the form of data for RESTful API
│   ├── registry                The skeleton of the application. Dependency injection is included. Connect interface with implementation
│   └── usecase                 Application specific business rules. Interfaces are included
│       ├── interactor              Implementation of the use case
│       └── model                   Data structs represents the input and output of the use cases
├── cmd                     Command line and executed program
│   ├── gentool                 Code generate tools
│   └── service                 The main executed program
├── deployments             Includes the configuration and the deployments for all environments
│   ├── heroku                  Heroku environment
│   ├── k8s                     K8s deployments
│   └── local                   Local develop environment
├── docs                    All documentations belong in here
├── internal                Internal library packages (not be exposed)
├── pkg                     Library packages
└── scripts                 Shell script files
```
