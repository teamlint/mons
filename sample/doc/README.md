## layout
### sample service 

./sample/
├── Taskfile.yml
├── adapter
│   ├── application
│   │   └── service
│   │       └── user_service.go
│   ├── domain
│   │   └── service
│   ├── repository
│   │   └── mysql
│   │       └── user_repository.go
│   └── ui
│       └── web
│           ├── handler
│           │   └── home.go
│           └── server
│               └── server.go
├── application
│   ├── command
│   │   ├── create_user_command.go
│   │   └── update_user_command.go
│   ├── dto
│   │   └── user.go
│   ├── endpoint
│   │   ├── middleware.go
│   │   └── user_endpoint.go
│   ├── query
│   │   └── user_query.go
│   ├── service
│   │   ├── middleware.go
│   │   ├── service.go
│   │   └── user_service.go
│   └── transport
│       ├── http
│       │   └── handler.go
│       ├── nats
│       │   └── handler.go
│       └── natshttp
│           └── handler.go
├── client
│   ├── http
│   │   └── client.go
│   ├── nats
│   │   └── client.go
│   └── natshttp
│       └── client.go
├── cmd
│   ├── client
│   │   ├── internal
│   │   │   └── internal.go
│   │   └── main.go
│   ├── server
│   │   └── main.go
│   └── ui
│       └── web
│           ├── main.go
│           └── views
│               ├── about
│               │   └── index.html
│               ├── home
│               │   ├── index.html
│               │   └── user.html
│               └── layouts
│                   ├── footer.html
│                   └── master.html
├── doc
│   └── README.md
├── domain
│   ├── model
│   │   └── user.go
│   ├── repository
│   │   └── user_repository.go
│   └── service
│       └── user_service.go
├── example
├── server
│   └── server.go
└── test


### shared package
./shared
├── adapter
│   └── repository
│       ├── gorm_repository_context.go
│       ├── gorm_transaction.go
│       └── mysql
│           └── gorm_translator.go
├── application
│   └── transaction
│       └── transaction.go
└── domain
    └── repository
        ├── repository_context.go
        ├── repository.go
        ├── transaction_scoper.go
        └── unit_of_work.go

## think
- service 接口参数及返回值对象和 endpoint 参数及返回值对象是否分隔
- application service 参数使用command/query还是使用endpoint xxxRequest/xxxResponse

## TODO
- config 
- CQRS
- middleware
- vendor
