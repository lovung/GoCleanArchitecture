# Golang Clean Architecture template

## Skeleton

.
├── app
│   ├── domain
│   │   ├── entity
│   │   └── repository
│   ├── external
│   │   ├── adaptor
│   │   ├── framework
│   │   └── persistence
│   │       ├── fs
│   │       ├── kv
│   │       ├── nosql
│   │       └── rdbms
│   ├── interface
│   │   ├── grpc
│   │   └── restful
│   │       ├── handler
│   │       ├── middleware
│   │       └── presenter
│   ├── registry
│   └── usecase
├── cmd
│   ├── gentool
│   └── service
├── deployments
│   ├── heroku
│   ├── k8s
│   └── local
├── docs
├── internal
├── pkg
└── scripts
