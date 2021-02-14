![Go](https://github.com/lovung/GoCleanArchitecture/workflows/Go/badge.svg)
![Go Report](https://goreportcard.com/badge/github.com/lovung/GoCleanArchitecture?style=flat-square)
[![codecov](https://codecov.io/gh/lovung/GoCleanArchitecture/branch/main/graph/badge.svg?token=QOX2GKGTA2)](https://codecov.io/gh/lovung/GoCleanArchitecture)

# Golang Clean Architecture template	
## Clean Architecture blog post	
[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)	
![Layers](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)	

## Skeleton
```
.
├── app
│   ├── config
│   ├── external
│   │   ├── adaptor
│   │   ├── api
│   │   └── persistence
│   │       ├── fs
│   │       ├── kv
│   │       ├── nosql
│   │       └── rdbms
│   ├── internal
│   │   ├── apperror
│   │   ├── domain
│   │   │   ├── entity
│   │   │   └── repository
│   │   │       └── mockrepo
│   │   ├── interface
│   │   │   ├── grpc
│   │   │   ├── persistence
│   │   │   │   ├── fs
│   │   │   │   ├── kv
│   │   │   │   ├── nosql
│   │   │   │   └── rdbms
│   │   │   │       └── gormrepo
│   │   │   └── restful
│   │   │       ├── handler
│   │   │       ├── middleware
│   │   │       └── presenter
│   │   ├── transaction
│   │   │   └── mocktrans
│   │   └── usecase
│   │       ├── dto
│   │       ├── interactor
│   │       └── mockusecase
│   └── registry
├── build
├── cmd
│   ├── gentool
│   └── services
│       └── core
├── databases
│   └── mysql
│       └── migrations
├── deployments
│   ├── heroku
│   ├── k8s
│   └── local
├── docs
├── pkg
│   ├── config
│   ├── copier
│   ├── gormutil
│   ├── hasher
│   ├── jwtutil
│   ├── logger
│   └── storage
├── scripts
└── test-results
```

